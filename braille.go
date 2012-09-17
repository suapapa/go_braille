// Copyright 2012, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*  Filename:    braille.go
 *  Author:      Homin Lee <homin.lee@suapapa.net>
 *  Created:     2012-08-25 02:54:03.999792 +0900 KST
 *  Description: Main source file in braille
 */

/*
  Package braille is a tool for unicode braille code.

  A braille character has 8 dots. 2 dots width and 4 dots height.
  Each dot called in following index.

    14
    25
    36
    78

  For example, A braille character which has dot 1, 5, 6, 7
  look like;

    ^_
    _^
    _^
    ^_

  This package provide functions for convert such a
  dot representation to unicode and vice versa

*/
package braille

import (
	"fmt"
	"log"
)

var BrailleCodeBase = rune(0x2800)

// Convert a unicode braille character to dot representation.
func Dot(r rune) []uint {
	if (r & 0xff00) != BrailleCodeBase {
		//log.Errorf("0x%x is not braille code~\n", r)
		fmt.Printf("0x%x is not braille code~\n", r)
		return nil
	}
	dots := make([]uint, 0)
	var i uint
	for i = 0; i < 8; i++ {
		if ((r >> i) & 1) == 1 {
			dots = append(dots, i+1)
		}
	}
	return dots
}

// Convert dot representation to a unicode character.
func Rune(dots ...uint) (r rune) {
	r = BrailleCodeBase
	for _, d := range dots {
		if 1 > d || d > 8 {
			fmt.Printf("Dot index must in range 1 to 8. Not %d\n", d)
			return 0x00
		}
		r |= 1 << (d - 1)
	}
	return
}

var MarkerNumber = Rune(3, 4, 5, 6)
var MarkerCap = Rune(6)

var number = map[int]rune{
	1: Rune(1),
	2: Rune(1, 2),
	3: Rune(1, 4),
	4: Rune(1, 4, 5),
	5: Rune(1, 5),
	6: Rune(1, 2, 4),
	7: Rune(1, 2, 4, 5),
	8: Rune(1, 2, 5),
	9: Rune(2, 4),
	0: Rune(2, 4, 5),
}

// Return braille code for given number and English alphabet
func Alphabet(c rune) (rs []rune) {
	if 'A' <= c && c <= 'Z' {
		rs = append(rs, MarkerCap)
		c ^= 0x20
	}

	var a rune

	switch {
	case '0' <= c && c <= '9':
		i := int(c - '0')
		a = number[i]

	case ('a' <= c && c <= 'v') || ('x' <= c && c <= 'z'):
		i := int(c - 'a')
		if 'a' <= c && c <= 'v' {
			i += 1
		}

		a = number[i%10]
		switch i / 10 {
		case 1:
			a |= Rune(3)
		case 2:
			a |= Rune(3, 6)
		}

	case c == 'w':
		a = number[0] | Rune(6)

	case c == ' ':
		a = BrailleCodeBase

	default:
		log.Printf("Braille for %c not present... yet\n", c)
		a = BrailleCodeBase
	}

	rs = append(rs, a)

	return
}

// Encode input alpha-numeric string to braille string
func Encode(s string) string {
	rs := make([]rune, len(s))

	var lastC rune
	for _, c := range s {
		if ('0' <= c && c <= '9') && ('a' <= lastC && lastC <= 'z') {
			rs = append(rs, 0x20, MarkerNumber)
		}
		rs = append(rs, Alphabet(c)...)

		if c != ' ' && c != '-' {
			lastC = c | 0x20
		}
	}

	return string(rs)
}
