// Copyright 2012, Homin Lee. All rights reserved.
// Use of this source braille.Code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*  Filename:    braille.go
 *  Author:      Homin Lee <homin.lee@suapapa.net>
 *  Created:     2012-08-25 02:54:03.999792 +0900 KST
 *  Description: Main source file in braille/ko
 */

package ko

import (
	"log"

	brl "github.com/suapapa/go_braille"
	han "github.com/suapapa/go_hangul"
)

/*
자음은 초성과 종성의 모양이 다르지만, 유사성을 띠고 있다.
종성의 모양은 초성의 좌우를 뒤집어 만들거나, 모양을 아래쪽으로 한 칸 내려서 만든다.
자음 점자는 두 줄 이하만을 사용한다.
초성 ㅇ은 점자에서 따로 쓰지 않으며,
별도의 기호가 있는 ㅆ를 뺀 겹받침은 해당하는 낱자를 순서대로 적는다.
*/
var jamo = map[rune][]rune{
	// 초성
	han.LeadG: {brl.Rune(4)},
	han.LeadN: {brl.Rune(1, 4)},
	han.LeadD: {brl.Rune(2, 4)},
	han.LeadR: {brl.Rune(5)},
	han.LeadM: {brl.Rune(1, 5)},
	han.LeadB: {brl.Rune(4, 5)},
	han.LeadS: {brl.Rune(6)},
	han.LeadJ: {brl.Rune(4, 6)},
	han.LeadC: {brl.Rune(5, 6)},
	han.LeadK: {brl.Rune(1, 2, 4)},
	han.LeadT: {brl.Rune(1, 2, 5)},
	han.LeadP: {brl.Rune(1, 4, 5)},
	han.LeadH: {brl.Rune(2, 4, 5)},

	// 종성
	han.TailG:  {brl.Rune(1)},
	han.TailN:  {brl.Rune(2, 4)},
	han.TailD:  {brl.Rune(3, 4)},
	han.TailL:  {brl.Rune(2)},
	han.TailM:  {brl.Rune(2, 6)},
	han.TailB:  {brl.Rune(1, 2)},
	han.TailS:  {brl.Rune(3)},
	han.TailSS: {brl.Rune(3, 4)},
	han.TailNG: {brl.Rune(2, 3, 5, 6)},
	han.TailJ:  {brl.Rune(1, 3)},
	han.TailC:  {brl.Rune(3, 3)},
	han.TailK:  {brl.Rune(2, 3, 5)},
	han.TailT:  {brl.Rune(2, 3, 6)},
	han.TailP:  {brl.Rune(2, 5, 6)},
	han.TailH:  {brl.Rune(3, 5, 6)},

	// 모음
	han.MedialA:   {brl.Rune(1, 2, 6)},
	han.MedialYA:  {brl.Rune(3, 4, 5)},
	han.MedialEO:  {brl.Rune(2, 3, 4)},
	han.MedialYEO: {brl.Rune(1, 5, 6)},
	han.MedialO:   {brl.Rune(1, 3, 6)},
	han.MedialYO:  {brl.Rune(3, 4, 6)},
	han.MedialU:   {brl.Rune(1, 3, 4)},
	han.MedialYU:  {brl.Rune(1, 4, 6)},
	han.MedialEU:  {brl.Rune(2, 4, 6)},
	han.MedialI:   {brl.Rune(1, 3, 5)},
	han.MedialAE:  {brl.Rune(1, 2, 3, 5)},
	han.MedialE:   {brl.Rune(1, 3, 4, 5)},

	// 이중모음
	han.MedialYAE: {brl.Rune(3, 4, 5), brl.Rune(1, 2, 3, 5)},
	han.MedialYE:  {brl.Rune(3, 4)},
	han.MedialWA:  {brl.Rune(1, 2, 3, 6)},
	han.MedialWAE: {brl.Rune(1, 2, 3, 6), brl.Rune(1, 2, 3, 5)},
	han.MedialOE:  {brl.Rune(1, 3, 4, 5, 6)},
	han.MedialWEO: {brl.Rune(1, 2, 3, 4)},
	han.MedialWE:  {brl.Rune(1, 2, 3, 4), brl.Rune(1, 2, 3, 5)},
	han.MedialWI:  {brl.Rune(1, 3, 4), brl.Rune(1, 2, 3, 5)},
	han.MedialYI:  {brl.Rune(2, 4, 5, 6)},
}

var markerDoubleLead = brl.Rune(6)          // 된소리표
var markerAcientRune = brl.Rune(5)          // 고문자표
var markerNumber = brl.Rune(3, 4, 5, 6)     // 수표
var markerJamo = brl.Rune(1, 2, 3, 4, 5, 6) // 말소표
var markerForeign = brl.Rune(3, 5, 6)       // 외국어표

// Jamo convert Hangul-jamo to corresponding braille
func Jamo(r rune) (rs []rune) {
	switch {
	case han.IsLead(r):
		if b, ok := jamo[r]; ok {
			rs = append(rs, b...)
		} else if bs, ok := han.SplitMultiElement(r); ok {
			if bs[0] == bs[1] {
				singleLead := han.Lead(bs[0])
				rs = append(rs, markerDoubleLead)
				rs = append(rs, jamo[singleLead]...)
			} else {
				log.Printf("No braille code for lead, %c -> %c + %c\n",
					han.CompatJamo(r), bs[0], bs[1])
			}
		} else {
			switch r {
			case han.LeadZS:
				// lead consonant can be omitted when it's ZS.
			default:
				log.Printf("No braille code for lead, %c\n",
					han.CompatJamo(r))
			}
		}

	case han.IsMedial(r):
		if b, ok := jamo[r]; ok {
			rs = append(rs, b...)
		} else {
			log.Printf("No braille code for vowel, %c\n",
				han.CompatJamo(r))
		}

	case han.IsTail(r):
		if b, ok := jamo[r]; ok {
			rs = append(rs, b...)
		} else if bs, ok := han.SplitMultiElement(r); ok {
			for _, c := range bs {
				rs = append(rs, jamo[han.Tail(c)]...)
			}
		} else {
			switch r {
			case 0:
				// tail consonant can be omitted.
			default:
				log.Printf("No braille code for tail, %c\n",
					han.CompatJamo(r))
			}
		}
	case r == 0:
		//log.Printf("Mayby, happened when no tail exists\n");
		return

	default:
		log.Printf("%c(0x%0x) seems not a Hangul-jamo!\n",
			han.CompatJamo(r), r)
		return nil
	}

	return rs
}

// 기타, 문장부호
var symbol = map[rune][]rune{
	'.': {brl.Rune(2, 5, 6)},
	',': {brl.Rune(5)},
	':': {brl.Rune(5), brl.Rune(2)},
	';': {brl.Rune(5, 6), brl.Rune(2, 3)},
	'?': {brl.Rune(2, 3, 6)},
	'!': {brl.Rune(4, 5, 6)},
	'…': {brl.Rune(5), brl.Rune(5), brl.Rune(5)},
	'‘': {brl.Rune(1, 2, 6)},
	'’': {brl.Rune(3, 5, 6)},
	'“': {brl.Rune(6), brl.Rune(2, 3, 6)},
	'”': {brl.Rune(3, 5, 6), brl.Rune(3)},
	'(': {brl.Rune(3, 6)},
	')': {brl.Rune(3, 6)},
	'{': {brl.Rune(2, 3, 6), brl.Rune(2, 3)},
	'}': {brl.Rune(5, 6), brl.Rune(3, 5, 6)},
	'[': {brl.Rune(2, 3, 6), brl.Rune(3)},
	']': {brl.Rune(6), brl.Rune(3, 5, 6)},
	'-': {brl.Rune(3, 6)},
	'~': {brl.Rune(3, 6), brl.Rune(3, 6)},
	'_': {brl.Rune(6), brl.Rune(3, 6), brl.Rune(3, 6), brl.Rune(3)},
	'/': {brl.Rune(4, 5, 6), brl.Rune(3, 4)},
	'※': {brl.Rune(3, 5), brl.Rune(3, 5)},
	'—': {brl.Rune(5, 6), brl.Rune(3, 6), brl.Rune(3, 6), brl.Rune(2, 3)}, // 줄표
	// TODO: 빠진 문장 부호 있음.
}

func isSupportedSymbol(c rune) bool {
	_, ok := symbol[c]
	return ok
}

/*
약자 :
주 쓰이는 글자나 단어에는 별도의 기호가 배당되어 있다.
특히 약어의 경우 첫 글자와 마지막 글자에서 한 낱자씩 따 온 것이다.
약자는 모호하지 않다면 앞뒤에 다른 낱자를 붙여서 쓸 수 있다.
(까 = ㄱ + 가, 말 = 마 + ㄹ 등)
*/
var abbr = map[rune][]rune{
	'가': {brl.Rune(1, 2, 4, 6)},
	'나': {brl.Rune(1, 4)},
	'다': {brl.Rune(2, 4)},
	'마': {brl.Rune(1, 5)},
	'바': {brl.Rune(4, 5)},
	'사': {brl.Rune(1, 2, 3)},
	'자': {brl.Rune(4, 6)},
	'카': {brl.Rune(1, 2, 4)},
	'타': {brl.Rune(1, 2, 5)},
	'파': {brl.Rune(1, 4, 5)},
	'하': {brl.Rune(2, 4, 5)},
	'것': {brl.Rune(4, 5, 6), brl.Rune(2, 3, 4)},
	'억': {brl.Rune(1, 4, 5, 6)},
	'언': {brl.Rune(2, 3, 4, 5, 6)},
	'얼': {brl.Rune(2, 3, 4, 5)},
	'연': {brl.Rune(1, 6)},
	'열': {brl.Rune(1, 2, 5, 6)},
	'엉': {brl.Rune(1, 2, 4, 5, 6)},
	'영': {brl.Rune(1, 2, 4, 5, 6)},
	'옥': {brl.Rune(1, 3, 4, 6)},
	'온': {brl.Rune(1, 2, 3, 5, 6)},
	'옹': {brl.Rune(1, 2, 3, 4, 5, 6)},
	'운': {brl.Rune(1, 2, 4, 5)},
	'울': {brl.Rune(1, 2, 3, 4, 6)},
	'은': {brl.Rune(1, 3, 5, 6)},
	'을': {brl.Rune(2, 3, 4, 6)},

	// TODO: 약어(예, 그래서) 빠짐
}

// TODO: 약자, 약어 처리를 위해서는 형태소 분석 필요! 일단 다 풀어 씀. :P

func hasHangul(s string) bool {
	for _, c := range s {
		if han.IsHangul(c) {
			return true
		}
	}
	log.Printf("%s not has hangul\n", s)
	return false
}

// Encode encodes Hangul string to braille
func Encode(s string) (string, int) {
	needForeignMarker := hasHangul(s)

	var currentMarker rune
	rs := make([]rune, 0)
	for _, c := range s {
		switch {
		case han.IsJaeum(c) || han.IsMoeum(c):
			if currentMarker != markerJamo {
				rs = append(rs, markerJamo)
				currentMarker = markerJamo
			}
			rs = append(rs, Jamo(c)...)
		case han.IsHangul(c):
			// In Korean braille string there is no marker for
			// Korean string.
			if currentMarker != 0x00 {
				currentMarker = 0x00
			}
			i, m, f := han.Split(c)
			rs = append(rs, Jamo(i)...)
			rs = append(rs, Jamo(m)...)
			rs = append(rs, Jamo(f)...)
		case '0' <= c && c <= '9':
			if currentMarker != markerNumber {
				rs = append(rs, markerNumber)
				currentMarker = markerNumber
			}
			rs = append(rs, brl.Alphabet(c)...)
		case ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z'):
			if currentMarker != markerForeign {
				if needForeignMarker {
					rs = append(rs, markerForeign)
				}
				currentMarker = markerForeign
			}
			rs = append(rs, brl.Alphabet(c)...)
		case isSupportedSymbol(c):
			// No marker change in symbols... Is it right?
			rs = append(rs, symbol[c]...)
		case c == 0x20:
			rs = append(rs, brl.Rune())
			currentMarker = 0x00
		case c == 0x0a:
			rs = append(rs, c)
			currentMarker = 0x00
		default:
			log.Printf("Braille for %c(0x%x) not present... yet", c, c)
			rs = append(rs, brl.BrailleCodeBase)
		}

	}

	return string(rs), len(rs)
}
