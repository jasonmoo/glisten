package main

import (
	"math"

	"github.com/fogleman/gg"
)

func main() {

	const (
		w = 1440
		h = 900
		r = h * math.Phi
	)

	dc := gg.NewContext(w, h)

	dc.SetRGB(0, 0, 0)
	dc.DrawRectangle(0, 0, w, h)
	dc.Fill()

	dc.SetLineCapRound()
	dc.Translate(w/2, h/2)

	rad := gg.Radians(200)

	dc.Scale(.5, .5)

	for a := float64(1); a < r; a += ((r - a) / 1000) + 0.06 {
		switch int(a) % 3 {
		case 0:
			dc.SetRGBA255(int(math.Mod(a, 200)+55), 255, 0, 2)
		case 1:
			dc.SetRGBA255(int(a/2), 255, 0, 4)
		case 2:
			dc.SetRGBA255(int(math.Mod(a, 255)), 255, 0, 4)
		}
		dc.SetLineWidth(math.Mod(a, 300))
		dc.MoveTo(0, 0)
		dc.QuadraticTo(math.Mod(a, 100), math.Mod(a, 200), 150, 50)
		dc.Rotate(rad)
		dc.Translate(a, a)
		dc.Stroke()
	}

	dc.SavePNG("out.png")

}
