// Copyright 2012, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package braille

/*  Filename:    braille_test.go
 *  Author:      Homin Lee <homin.lee@suapapa.net>
 *  Created:     2012-08-25 02:54:04.000269 +0900 KST
 *  Description: Main test file for braille
 */

import (
	"fmt"
	"testing"
)

func TestDot(t *testing.T) {
	fmt.Println("0x28C1 =", Dot(0x28c1))
	fmt.Println("0x282D =", Dot(0x282D))
	fmt.Println("0x28BF =", Dot(0x28BF))
	fmt.Println("0x28FF =", Dot(0x28FF))

}

func TestCode(t *testing.T) {
	fmt.Printf("1-5-6-7 -> %c\n", Rune(1, 5, 6, 7))
	fmt.Printf("1-2-3-4-5-6-7-8 -> %c\n", Rune(1, 2, 3, 4, 5, 6, 7, 8))
}

func TestNumber(t *testing.T) {
	for _, c := range "1234567890" {
		fmt.Printf("%c : %c\n", c, Alphabet(c))
	}
}

func TestAlphabet(t *testing.T) {
	s := "HackTime for Google Hackfair 2012-09-01"
	fmt.Println(s)
	fmt.Println(Encode(s))
}
