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

	f, err := os.Create("test.svg")
	if err != nil {
		t.Errorf("failed to create test.svg : %s\n", err)
	}
	defer f.Close()

	canvas := svg.New(f)
	defer canvas.End()

	dot := 10
	margin := 3

	s := "Braille Printer"
	bs, bsc := brl.Encode(s)
	fmt.Println(bs)

	cw := bsc*(dot*2) + (bsc+1)*margin
	ch := margin*2 + dot*4

	canvas.Start(cw, ch)

	x := margin
	for _, c := range bs {
		if c & 0x2800 != 0x2800 {
			continue
		}
		Draw(canvas, c, x, margin, dot)
		x += dot * 2
		x += margin
	}
}
