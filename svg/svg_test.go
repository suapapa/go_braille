// Copyright 2012, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package svg

import (
	"fmt"
	svg "github.com/ajstarks/svgo"
	brl "github.com/suapapa/go_braille"
	"os"
	"testing"
)

func TestDraw(t *testing.T) {
	f, err := os.Create("test_label.svg")
	if err != nil {
		t.Errorf("failed to create test.svg : %s\n", err)
	}
	defer f.Close()

	canvas := svg.New(f)
	defer canvas.End()

	s := "Braille Printer"
	bs, _ := brl.Encode(s)
	DrawLabel(canvas, bs)
}

func TestCalcCanvasSize_30x27(t *testing.T) {
	f, err := os.Create("test_page.svg")
	if err != nil {
		t.Errorf("failed to create test.svg : %s\n", err)
	}
	defer f.Close()

	canvas := svg.New(f)
	defer canvas.End()

	br := make([]rune, 100, 0x28FF)
	for i := 0; i < 100; i++ {
		switch {
		case i%5 == 0:
			br[i] = 0x2800
		case i%37 == 0:
			br[i] = 0xa
		default:
			br[i] = brl.Rune(1,2,3,4,5,6)
		}
	}
	bs := string(br)
	fmt.Println(bs)
	calcLines(bs, 30)

	DrawPage30(canvas, bs)
}
