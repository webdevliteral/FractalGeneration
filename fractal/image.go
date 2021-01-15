package fractal

import (
	"image"
	"image/color"
	"sync"
)

func GenFractalImage(fractal Fractal, imageWidth int, imageHeight int, numBands int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	bandHeight := imageHeight / numBands

	var wg sync.WaitGroup
	for i := 0; i < numBands; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			start := i * bandHeight
			for y := start; y < start+bandHeight; y++ {
				for x := 0; x < imageWidth; x++ {
					xComplex := float64(x-imageWidth/2) / float64(imageWidth) * 4
					yComplex := float64(y-imageHeight/2) / float64(imageHeight) * 4

					c := complex(xComplex, yComplex)

					gradient := fractal.FractalFunc(c)
					col := uint8(gradient * 255)

					img.Set(x, y, color.RGBA{col, col, 0, 255})

				}
			}
		}(i)
	}

	wg.Wait()

	return img
}
