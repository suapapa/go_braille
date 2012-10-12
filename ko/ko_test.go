// Copyright 2012, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ko

/*  Filename:    ko_test.go
 *  Author:      Homin Lee <homin.lee@suapapa.net>
 *  Created:     2012-08-25 02:54:04.000269 +0900 KST
 *  Description: Main test file for braille/ko
 */

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestHangul(t *testing.T) {
	s := "오빤 강남 스타일"
	fmt.Println(s)
	fmt.Println(Encode(s))
}

func TestHangulMultiline(t *testing.T) {
	s := "낮에는 따사로운 인간적인 여자\n밤이오면 심장이 뜨거워 지는 여자"
	fmt.Println(s)
	bs, bsLen := Encode(s)
	rLen := utf8.RuneCountInString(bs)
	fmt.Println(bs, bsLen, rLen)
}
