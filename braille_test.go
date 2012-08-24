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

func TestDots(t *testing.T) {
	fmt.Println("0x28C1 =", Dots(0x28c1))
	fmt.Println("0x282D =", Dots(0x282D))
	fmt.Println("0x28BF =", Dots(0x28BF))
	fmt.Println("0x28FF =", Dots(0x28FF))

}

func TestCode(t *testing.T) {
	fmt.Printf("%c\n", Code(&[]uint32{1, 3, 4, 6}))
	fmt.Printf("%c\n", Code(&[]uint32{1, 2, 3, 4, 5, 6, 7, 8}))
}
