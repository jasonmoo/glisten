package main

import (
	"math"

	"github.com/fogleman/gg"
)

func main() {

	// orig javascript implementation
	//
	// b.lineCap = "round";
	// for (var a = 1;a < 600; a += .4) {
	//   b.beginPath();
	//   b.strokeStyle = [
	//  	"rgba(" + ((a | 0) % 200 + 55) + ",255,0,.01)",
	//  	"rgba(" + (a / 2 | 0) + ",255,0,.02)",
	//  	"rgba(" + (a | 0) % 255 + ",255,0,.02)",
	//  ][a % 3]
	//   b.lineWidth = a % 300;
	//   b.bezierCurveTo(0, 0, a % 100, a % 200, 150, 50);
	//   b.rotate(rad(200));
	//   b.translate(a, a);
	//   b.stroke();
	// }

	const (
		r = 600
		S = 970 // r * math.Phi
	)

	dc := gg.NewContext(S, S)

	dc.SetRGB(0, 0, 0)
	dc.DrawRectangle(0, 0, S, S)
	dc.Fill()

	dc.SetLineCapRound()
	dc.Translate(S/2, S/2)

	rad := gg.Radians(200)

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
