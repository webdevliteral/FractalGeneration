package fractal

import "math"

type Fractal interface {
	FractalFunc(c complex128) float64
}

func absComplex(c complex128) float64 {
	return math.Sqrt(real(c)*real(c) + imag(c)*imag(c))
}
