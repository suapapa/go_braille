// Copyright 2012, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package svg

import (
	svg "github.com/ajstarks/svgo"
	brl "github.com/suapapa/go_braille"
	"log"
)

func drawDot(c *svg.SVG, x, y, s int) {
	c.Rect(x, y, s, s, "fill:none; stroke:yellow")
	c.Circle(x+s/2, y+s/2, s/3)
}

// Draw braille rune r at left-top x,y, with dotsize s. Note! single dot's size
// so actual width of a braille character is 2*s and heghit is 4*s
func Draw(c *svg.SVG, r rune, x, y, s int) {
	if r&brl.BrailleCodeBase == 0 {
		log.Printf("%c is not braille code!\n", r)
		return
	}

	c.Rect(x, y, 2*s, 4*s, "fill:none; stroke:red")
	log.Println(brl.Dot(r))

	for _, di := range brl.Dot(r) {
		switch di {
		case 1:
			drawDot(c, x, y, s)
		case 2:
			drawDot(c, x, y+s, s)
		case 3:
			drawDot(c, x, y+2*s, s)
		case 4:
			drawDot(c, x+s, y, s)
		case 5:
			drawDot(c, x+s, y+s, s)
		case 6:
			drawDot(c, x+s, y+2*s, s)
		case 7:
			drawDot(c, x, y+3*s, s)
		case 8:
			drawDot(c, x+s, y+3*s, s)
		}
	}
}
