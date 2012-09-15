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
	"github.com/suapapa/go_braille"
	"github.com/suapapa/go_hangul"
	"io"
	"unicode/utf8"
)

var jaMulti = braille.Code([]uint{6}) // 된소리표

var ko = map[rune]rune{
	// 초성
	hangul.LEAD_G: braille.Code([]uint{4}),
	hangul.LEAD_N: braille.Code([]uint{1, 4}),
	hangul.LEAD_D: braille.Code([]uint{2, 4}),
	hangul.LEAD_R: braille.Code([]uint{5}),
	hangul.LEAD_M: braille.Code([]uint{1, 5}),
	hangul.LEAD_B: braille.Code([]uint{4, 5}),
	hangul.LEAD_S: braille.Code([]uint{6}),
	hangul.LEAD_J: braille.Code([]uint{4, 6}),
	hangul.LEAD_C: braille.Code([]uint{5, 6}),
	hangul.LEAD_K: braille.Code([]uint{1, 2, 4}),
	hangul.LEAD_T: braille.Code([]uint{1, 2, 5}),
	hangul.LEAD_P: braille.Code([]uint{1, 4, 5}),
	hangul.LEAD_H: braille.Code([]uint{2, 4, 5}),

	// 종성
	hangul.TAIL_G:  braille.Code([]uint{1}),
	hangul.TAIL_N:  braille.Code([]uint{2, 4}),
	hangul.TAIL_D:  braille.Code([]uint{3, 4}),
	hangul.TAIL_L:  braille.Code([]uint{2}),
	hangul.TAIL_M:  braille.Code([]uint{2, 6}),
	hangul.TAIL_B:  braille.Code([]uint{1, 2}),
	hangul.TAIL_S:  braille.Code([]uint{3}),
	hangul.TAIL_SS: braille.Code([]uint{3, 4}),
	hangul.TAIL_NG: braille.Code([]uint{2, 3, 5, 6}),
	hangul.TAIL_J:  braille.Code([]uint{1, 3}),
	hangul.TAIL_C:  braille.Code([]uint{3, 3}),
	hangul.TAIL_K:  braille.Code([]uint{2, 3, 5}),
	hangul.TAIL_T:  braille.Code([]uint{2, 3, 6}),
	hangul.TAIL_P:  braille.Code([]uint{2, 5, 6}),
	hangul.TAIL_H:  braille.Code([]uint{3, 5, 6}),

	// 모음
	hangul.MEDIAL_A:   braille.Code([]uint{1, 2, 6}),
	hangul.MEDIAL_YA:  braille.Code([]uint{3, 4, 5}),
	hangul.MEDIAL_EO:  braille.Code([]uint{2, 3, 4}),
	hangul.MEDIAL_YEO: braille.Code([]uint{1, 5, 6}),
	hangul.MEDIAL_O:   braille.Code([]uint{1, 3, 6}),
	hangul.MEDIAL_YO:  braille.Code([]uint{3, 4, 6}),
	hangul.MEDIAL_U:   braille.Code([]uint{1, 3, 4}),
	hangul.MEDIAL_YU:  braille.Code([]uint{1, 4, 6}),
	hangul.MEDIAL_EU:  braille.Code([]uint{2, 4, 6}),
	hangul.MEDIAL_I:   braille.Code([]uint{1, 3, 5}),
	hangul.MEDIAL_AE:  braille.Code([]uint{1, 2, 3, 5}),
	hangul.MEDIAL_E:   braille.Code([]uint{1, 3, 4, 5}),

	// 약자
	'가': braille.Code([]uint{1, 2, 4, 6}),
	'나': braille.Code([]uint{1, 4}),
	'다': braille.Code([]uint{2, 4}),
}

// http://blog.daum.net/wwwhangulo/5285748

/*
자음은 초성과 종성의 모양이 다르지만, 유사성을 띠고 있다. 종성의 모양은 초성의 좌우를 뒤집어 만들거나, 모양을 아래쪽으로 한 칸 내려서 만든다.
자음 점자는 두 줄 이하만을 사용한다. 초성 ㅇ은 점자에서 따로 쓰지 않으며, 별도의 기호가 있는 ㅆ를 뺀 겹받침은 해당하는 낱자를 순서대로 적는다.
*/

func runeWriter(r rune, w io.Writer) int {
	p := make([]byte, 3)
	c := utf8.EncodeRune(p, r)
	return c
}

func decode(h string, w io.Writer) {
	for _, c := range h {
		i, m, f := hangul.Split(c)
		runeWriter(i, w)
		runeWriter(m, w)
		runeWriter(f, w)
	}
}
