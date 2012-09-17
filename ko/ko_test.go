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
)

func TestHangul(t *testing.T) {
	s := "오빤 강남 스타일"
	fmt.Println(s)
	fmt.Println(Encode(s))
}
