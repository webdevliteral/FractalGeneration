package fractal

import "math"

type Julia struct {
	MaxIterations int
	Z             complex128
}

func (j *Julia) FractalFunc(c complex128) float64 {
	for i := 0; i < j.MaxIterations; i++ {
		if absComplex(c) > 2 {
			return (float64(i) + 1 - math.Log2(math.Log(absComplex(c)))) / float64(j.MaxIterations)
		}

		c = c*c + j.Z
	}

	return 0
}
