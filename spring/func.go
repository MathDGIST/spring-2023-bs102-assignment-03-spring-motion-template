package spring

import (
	"image"
	"image/color"
)

type Func struct {
	Formula func(float64) float64
	Domain  [2]float64
}

func NewFunc(f func(float64) float64, d [2]float64) *Func {
	return &Func{
		Formula: f,
		Domain:  d,
	}
}

// Graph graphs the function f on the image img with the given width and height.
// The f has its own domain.
// The bounds for the x and y axes of the canvas are given by xb and yb.
func (f Func) Graph(w, h int, xb, yb [2]float64) *image.Paletted {
	img := image.NewPaletted(
		image.Rect(0, 0, w, h),
		color.Palette{
			color.White,
			color.Black,
		})
	/* Insert your code below */

	/* ---------------------- */
	return img
}

// line returns the array of the pixel points on the straight line
// between two pixel coordinates (x1, y1) and (x2, y2).
func line(x1, y1, x2, y2 int) [][2]int {
	abs := func(r int) int {
		if r >= 0 {
			return r
		}
		return -r
	}
	var axisReversed bool
	if abs(x2-x1) < abs(y2-y1) {
		x1, y1 = y1, x1
		x2, y2 = y2, x2
		axisReversed = true
	}
	slope := float64(y2-y1) / float64(x2-x1)
	inc := 1
	if x2 < x1 {
		inc = -1
	}
	var arr [][2]int
	for i := 0; i <= abs(x2-x1); i++ {
		x, y := x1+i*inc, y1+int(float64(i*inc)*slope)
		if axisReversed {
			arr = append(arr, [2]int{y, x})
		} else {
			arr = append(arr, [2]int{x, y})
		}
	}
	return arr
}
