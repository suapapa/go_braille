// Copyright 2012, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package svg

import (
	svg "github.com/ajstarks/svgo"
	"os"
	"testing"
)

func TestDraw(t *testing.T) {
	canvas := svg.New(os.Stdout)
	canvas.Start(500, 500)

	//drawDot(canvas, 250, 250, 100)
	Draw(canvas, 0x285F, 50, 50, 60)

	canvas.End()
}
