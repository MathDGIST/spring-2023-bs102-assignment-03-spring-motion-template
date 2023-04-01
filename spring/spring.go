package spring

import (
	"fmt"
	"image/gif"
	"math"
	"os"
)

type Spring struct {
	M, B, K, X0, X1, T float64
	Solution           func(float64) float64
}

func NewSpring() *Spring {
	return &Spring{}
}

func (sp *Spring) HandleInputs() {
	imsg := []string{
		"\n",
		"====================================\n",
		"This program animates the solution \n",
		"of the spring motion modeled by \n",
		"the following differential equation \n",
		"on the interval [0,T].\n",
		"mx'' + bx' + kx = 0, \n",
		"x(0) = x0, x'(0) = x1\n",
		"====================================\n",
		"\n",
	}
	for _, m := range imsg {
		fmt.Printf("%s", m)
	}
	msg := []string{
		"m: ",
		"b: ",
		"k: ",
		"x0: ",
		"x1: ",
		"T: ",
	}
	var v [6]float64
	for i, m := range msg {
		fmt.Printf("Enter the value of %s", m)
		fmt.Scanf("%f\n", &v[i])
	}
	sp.M, sp.B, sp.K = v[0], v[1], v[2]
	sp.X0, sp.X1, sp.T = v[3], v[4], v[5]
}

// FindSolution finds the solution of the spring motion.
// The solution is stored in the field Solution.
func (sp *Spring) FindSolution() {
	d := sp.B*sp.B - 4*sp.M*sp.K
	switch {
	case d > 0: // When the characteristic polynomial has two real roots.

	case d == 0: // When the characteristic polynomial has one real root.

	case d < 0: // When the characteristic polynomial has two complex roots.

	}
}

func (sp *Spring) Animate(w, h, dur int, fn string) {
	frames := 60 * dur
	xb := [2]float64{0, sp.T}
	yb := [2]float64{math.MaxFloat64, -math.MaxFloat64}
	for i := 0; i <= 100; i++ {
		x := sp.T * float64(i) / 100
		y := sp.Solution(x)
		if y < yb[0] {
			yb[0] = y
		}
		if y > yb[1] {
			yb[1] = y
		}
	}
	g := &gif.GIF{
		LoopCount: 1,
	}
	for i := 0; i <= frames; i++ {
		f := NewFunc(sp.Solution, [2]float64{0, sp.T * float64(i) / float64(frames)})
		img := f.Graph(w, h, xb, yb)
		g.Image = append(g.Image, img)
		g.Delay = append(g.Delay, 1)
	}
	fp, _ := os.Create(fn)
	defer fp.Close()
	gif.EncodeAll(fp, g)
}
