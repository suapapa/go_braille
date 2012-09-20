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
	brl "github.com/suapapa/go_braille"
	han "github.com/suapapa/go_hangul"
	"log"
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
	han.LEAD_G: []rune{brl.Rune(4)},
	han.LEAD_N: []rune{brl.Rune(1, 4)},
	han.LEAD_D: []rune{brl.Rune(2, 4)},
	han.LEAD_R: []rune{brl.Rune(5)},
	han.LEAD_M: []rune{brl.Rune(1, 5)},
	han.LEAD_B: []rune{brl.Rune(4, 5)},
	han.LEAD_S: []rune{brl.Rune(6)},
	han.LEAD_J: []rune{brl.Rune(4, 6)},
	han.LEAD_C: []rune{brl.Rune(5, 6)},
	han.LEAD_K: []rune{brl.Rune(1, 2, 4)},
	han.LEAD_T: []rune{brl.Rune(1, 2, 5)},
	han.LEAD_P: []rune{brl.Rune(1, 4, 5)},
	han.LEAD_H: []rune{brl.Rune(2, 4, 5)},

	// 종성
	han.TAIL_G:  []rune{brl.Rune(1)},
	han.TAIL_N:  []rune{brl.Rune(2, 4)},
	han.TAIL_D:  []rune{brl.Rune(3, 4)},
	han.TAIL_L:  []rune{brl.Rune(2)},
	han.TAIL_M:  []rune{brl.Rune(2, 6)},
	han.TAIL_B:  []rune{brl.Rune(1, 2)},
	han.TAIL_S:  []rune{brl.Rune(3)},
	han.TAIL_SS: []rune{brl.Rune(3, 4)},
	han.TAIL_NG: []rune{brl.Rune(2, 3, 5, 6)},
	han.TAIL_J:  []rune{brl.Rune(1, 3)},
	han.TAIL_C:  []rune{brl.Rune(3, 3)},
	han.TAIL_K:  []rune{brl.Rune(2, 3, 5)},
	han.TAIL_T:  []rune{brl.Rune(2, 3, 6)},
	han.TAIL_P:  []rune{brl.Rune(2, 5, 6)},
	han.TAIL_H:  []rune{brl.Rune(3, 5, 6)},

	// 모음
	han.MEDIAL_A:   []rune{brl.Rune(1, 2, 6)},
	han.MEDIAL_YA:  []rune{brl.Rune(3, 4, 5)},
	han.MEDIAL_EO:  []rune{brl.Rune(2, 3, 4)},
	han.MEDIAL_YEO: []rune{brl.Rune(1, 5, 6)},
	han.MEDIAL_O:   []rune{brl.Rune(1, 3, 6)},
	han.MEDIAL_YO:  []rune{brl.Rune(3, 4, 6)},
	han.MEDIAL_U:   []rune{brl.Rune(1, 3, 4)},
	han.MEDIAL_YU:  []rune{brl.Rune(1, 4, 6)},
	han.MEDIAL_EU:  []rune{brl.Rune(2, 4, 6)},
	han.MEDIAL_I:   []rune{brl.Rune(1, 3, 5)},
	han.MEDIAL_AE:  []rune{brl.Rune(1, 2, 3, 5)},
	han.MEDIAL_E:   []rune{brl.Rune(1, 3, 4, 5)},

	// 이중모음
	han.MEDIAL_YAE: []rune{brl.Rune(3, 4, 5), brl.Rune(1, 2, 3, 5)},
	han.MEDIAL_YE:  []rune{brl.Rune(3, 4)},
	han.MEDIAL_WA:  []rune{brl.Rune(1, 2, 3, 6)},
	han.MEDIAL_WAE: []rune{brl.Rune(1, 2, 3, 6), brl.Rune(1, 2, 3, 5)},
	han.MEDIAL_OE:  []rune{brl.Rune(1, 3, 4, 5, 6)},
	han.MEDIAL_WEO: []rune{brl.Rune(1, 2, 3, 4)},
	han.MEDIAL_WE:  []rune{brl.Rune(1, 2, 3, 4), brl.Rune(1, 2, 3, 5)},
	han.MEDIAL_WI:  []rune{brl.Rune(1, 3, 4), brl.Rune(1, 2, 3, 5)},
	han.MEDIAL_YI:  []rune{brl.Rune(2, 4, 5, 6)},
}

var markerDoubleLead = brl.Rune(6)          // 된소리표
var markerAcientRune = brl.Rune(5)          // 고문자표
var markerNumber = brl.Rune(3, 4, 5, 6)     // 수표
var markerJamo = brl.Rune(1, 2, 3, 4, 5, 6) // 말소표
var markerForeign = brl.Rune(3, 5, 6)       // 외국어표

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
			case han.LEAD_ZS:
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
	'.': []rune{brl.Rune(2, 5, 6)},
	',': []rune{brl.Rune(5)},
	':': []rune{brl.Rune(5), brl.Rune(2)},
	';': []rune{brl.Rune(5, 6), brl.Rune(2, 3)},
	'?': []rune{brl.Rune(2, 3, 6)},
	'!': []rune{brl.Rune(4, 5, 6)},
	'…': []rune{brl.Rune(5), brl.Rune(5), brl.Rune(5)},
	'‘': []rune{brl.Rune(1, 2, 6)},
	'’': []rune{brl.Rune(3, 5, 6)},
	'“': []rune{brl.Rune(6), brl.Rune(2, 3, 6)},
	'”': []rune{brl.Rune(3, 5, 6), brl.Rune(3)},
	'(': []rune{brl.Rune(3, 6)},
	')': []rune{brl.Rune(3, 6)},
	'{': []rune{brl.Rune(2, 3, 6), brl.Rune(2, 3)},
	'}': []rune{brl.Rune(5, 6), brl.Rune(3, 5, 6)},
	'[': []rune{brl.Rune(2, 3, 6), brl.Rune(3)},
	']': []rune{brl.Rune(6), brl.Rune(3, 5, 6)},
	'-': []rune{brl.Rune(3, 6)},
	'~': []rune{brl.Rune(3, 6), brl.Rune(3, 6)},
	'_': []rune{brl.Rune(6), brl.Rune(3, 6), brl.Rune(3, 6), brl.Rune(3)},
	'/': []rune{brl.Rune(4, 5, 6), brl.Rune(3, 4)},
	'※': []rune{brl.Rune(3, 5), brl.Rune(3, 5)},
	'—': []rune{brl.Rune(5, 6), brl.Rune(3, 6), brl.Rune(3, 6), brl.Rune(2, 3)}, // 줄표
	// TODO: 빠진 문장 부호 있음.
}

/*
약자 :
주 쓰이는 글자나 단어에는 별도의 기호가 배당되어 있다.
특히 약어의 경우 첫 글자와 마지막 글자에서 한 낱자씩 따 온 것이다.
약자는 모호하지 않다면 앞뒤에 다른 낱자를 붙여서 쓸 수 있다.
(까 = ㄱ + 가, 말 = 마 + ㄹ 등)
*/
var abbr = map[rune][]rune{
	'가': []rune{brl.Rune(1, 2, 4, 6)},
	'나': []rune{brl.Rune(1, 4)},
	'다': []rune{brl.Rune(2, 4)},
	'마': []rune{brl.Rune(1, 5)},
	'바': []rune{brl.Rune(4, 5)},
	'사': []rune{brl.Rune(1, 2, 3)},
	'자': []rune{brl.Rune(4, 6)},
	'카': []rune{brl.Rune(1, 2, 4)},
	'타': []rune{brl.Rune(1, 2, 5)},
	'파': []rune{brl.Rune(1, 4, 5)},
	'하': []rune{brl.Rune(2, 4, 5)},
	'것': []rune{brl.Rune(4, 5, 6), brl.Rune(2, 3, 4)},
	'억': []rune{brl.Rune(1, 4, 5, 6)},
	'언': []rune{brl.Rune(2, 3, 4, 5, 6)},
	'얼': []rune{brl.Rune(2, 3, 4, 5)},
	'연': []rune{brl.Rune(1, 6)},
	'열': []rune{brl.Rune(1, 2, 5, 6)},
	'엉': []rune{brl.Rune(1, 2, 4, 5, 6)},
	'영': []rune{brl.Rune(1, 2, 4, 5, 6)},
	'옥': []rune{brl.Rune(1, 3, 4, 6)},
	'온': []rune{brl.Rune(1, 2, 3, 5, 6)},
	'옹': []rune{brl.Rune(1, 2, 3, 4, 5, 6)},
	'운': []rune{brl.Rune(1, 2, 4, 5)},
	'울': []rune{brl.Rune(1, 2, 3, 4, 6)},
	'은': []rune{brl.Rune(1, 3, 5, 6)},
	'을': []rune{brl.Rune(2, 3, 4, 6)},

	// TODO: 약어(예, 그래서) 빠짐
}

// TODO: 약자, 약어 처리를 위해서는 형태소 분석 필요! 일단 다 풀어 씀. :P

func Encode(s string) (string, int) {
	rs := make([]rune, 0)
	for _, c := range s {
		switch {
		case han.IsJaeum(c) || han.IsMoeum(c):
			rs = append(rs, markerJamo)
			rs = append(rs, Jamo(c)...)
		case han.IsHangul(c):
			i, m, f := han.Split(c)
			rs = append(rs, Jamo(i)...)
			rs = append(rs, Jamo(m)...)
			rs = append(rs, Jamo(f)...)
		default:
			log.Printf("Braille for %c not present... yet", c)
			rs = append(rs, brl.BrailleCodeBase)
		}

	}

	return string(rs), len(rs)
}
