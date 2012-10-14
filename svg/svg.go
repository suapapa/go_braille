// Copyright 2012, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package svg

import (
	svg "github.com/ajstarks/svgo"
	brl "github.com/suapapa/go_braille"
	"log"
	"strings"
	"unicode/utf8"
)

func drawDot(c *svg.SVG, x, y, s int) {
	//c.Rect(x, y, s, s, "fill:none; stroke:yellow")
	c.Circle(x+s/2, y+s/2, s/3)
}

// Draw braille rune r at left-top x,y, with dotsize s. Note! single dot's size
// so actual width of a braille character is 2*s and heghit is 4*s
func Draw(c *svg.SVG, r rune, x, y, s int) {
	if r&brl.BrailleCodeBase == 0 {
		log.Printf("%c is not braille code!\n", r)
		return
	}

	//c.Rect(x, y, 2*s, 4*s, "fill:none; stroke:red")

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

// Draw svg canvas size for multi line braille string.
// According to [Full Page Slate](http://goo.gl/m4DhD),
// A braille page has 30 cells x 27 lines and only use 6 dots for a cell
func calcLines(s string, lineCellsCnt int) (lineCnt int) {
	lines := strings.Split(s, "\n")
	lineCnt = len(lines)
	log.Printf("calcLines: string has %d lines\n", lineCnt)

	for _, line := range lines {
		// TODO: concern word wrappeing
		lineCnt += (utf8.RuneCountInString(line) / lineCellsCnt)
	}
	log.Printf("calcLines: %d lines when %d cells for a line\n",
		lineCnt, lineCellsCnt)

	return
}

type BrailleCanvas struct {
	lineCellsCnt int
	dotSize int // recommend multiple of 6 for a dot size
	pageMarginTop, pageMarginBottom, pageMarginLeft, pageMarginRight int
	gapCell, gapLine int
	canvasW, canvasH int
}

func DrawLabel(s *svg.SVG, bs string) *BrailleCanvas {
	var c BrailleCanvas

	// TODO: they hard coded
	c.dotSize = 12
	c.pageMarginTop = 3
	c.pageMarginBottom = 3
	c.pageMarginLeft = 3
	c.pageMarginRight = 3
	c.gapCell = 3

	c.canvasW = c.pageMarginLeft + c.pageMarginRight
	c.canvasW += c.dotSize * 2 * utf8.RuneCountInString(bs)
	c.canvasW += c.gapCell * utf8.RuneCountInString(bs) - 1

	c.canvasH = c.pageMarginTop + c.pageMarginBottom
	c.canvasH += c.dotSize * 3

	s.Start(c.canvasW, c.canvasH)

	x := c.pageMarginLeft
	y := c.pageMarginTop
	for _, b := range bs {
		if (b&0x2800 != 0x2800) {
			log.Printf("0x%x is not a braille character!\n", b)
			continue
		}
		Draw(s, b, x, y, c.dotSize)
		x += (c.dotSize * 2) + c.gapCell
	}

	return &c
}

func DrawPage30(s *svg.SVG, bs string) *BrailleCanvas {
	var c BrailleCanvas

	// TODO: they hard coded
	c.dotSize = 12
	c.pageMarginTop = 10
	c.pageMarginBottom = 10
	c.pageMarginLeft = 10
	c.pageMarginRight = 10
	c.gapCell = 3
	c.gapLine = 6
	c.lineCellsCnt = 30

	c.canvasW = c.pageMarginLeft + c.pageMarginRight
	c.canvasW += c.dotSize * 2 * 30
	c.canvasW += c.gapCell * 30 - 1

	lineCnt := calcLines(bs, 30)
	c.canvasH = c.pageMarginTop + c.pageMarginBottom
	c.canvasH += c.dotSize * 3 * lineCnt
	c.canvasH += c.gapLine * lineCnt - 1

	s.Start(c.canvasW, c.canvasH)

	x := c.pageMarginLeft
	y := c.pageMarginTop

	lines := strings.Split(bs, "\n")
	ri := 0
	for _, line := range lines {
		for _, b := range line {
			if ri == 30 {
				x = c.pageMarginLeft
				y += (c.dotSize * 3) + c.gapLine
				ri = 0
			}

			if (b&0x2800 != 0x2800) {
				log.Printf("0x%x is not a braille character!\n", b)
				continue
			}

			Draw(s, b, x, y, c.dotSize)
			ri += 1
			x += (c.dotSize * 2) + c.gapCell
		}
		x = c.pageMarginLeft
		y += (c.dotSize * 3) + c.gapLine
		ri = 0
	}

	return &c
}

