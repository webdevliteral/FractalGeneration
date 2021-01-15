package fractal

import "math"

type Mandelbrot struct {
	MaxIterations int
}

func (m *Mandelbrot) FractalFunc(c complex128) float64 {
	var z complex128

	for i := 0; i < m.MaxIterations; i++ {
		if absComplex(z) > 2 {
			return (float64(i) + 1 - math.Log2(math.Log(absComplex(z)))) / float64(m.MaxIterations)
		}

		z = z*z + c
	}

	return 0
}
