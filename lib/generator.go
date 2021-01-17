package lib

import (
	"image"
	"image/color"
	"sync"
)

// PointFunc is an integrated fractal & color function used by a generator
type PointFunc func(xCoord, yCoord float64, iterationCap int) (R, G, B, A float64)

// Generator is our runner!
type Generator struct {
	Img          *image.NRGBA
	xPos         float64
	yPos         float64
	zoom         float64
	scaler       float64
	width        int
	height       int
	routines     int
	iterationCap int
	fn           PointFunc
	samples      int
}

// NewGenerator returns a generator!
func NewGenerator(width, height, routines, iterationCap, samples int, xPos, yPos, zoom float64, fn PointFunc) Generator {

	// Pick the smaller of the two dimensions (width and height) and use that length in
	// pixels as the length of 2 divided by the zoom factor as the scale for both axis.
	var scaler float64
	if width < height {
		scaler = float64(width)
	} else {
		scaler = float64(height)
	}

	return Generator{
		image.NewNRGBA(image.Rect(0, 0, width, height)),
		xPos,
		yPos,
		zoom,
		scaler,
		width,
		height,
		routines,
		iterationCap,
		fn,
		samples,
	}
}

// Generate spins out our workers!
func (f Generator) Generate() {
	var wg sync.WaitGroup
	wg.Add(f.routines)

	for routine := 0; routine < f.routines; routine++ {
		go f.genRoutine(&wg, routine)
	}

	wg.Wait()
}

func (f Generator) pixToCoord(xPix, yPix float64) (xCoord, yCoord float64) {
	xCoord = ((xPix - (float64(f.width) / 2)) * ((2 / f.scaler) / f.zoom)) + f.xPos
	yCoord = ((yPix - (float64(f.height) / 2)) * ((2 / f.scaler) / f.zoom)) + f.yPos
	return xCoord, yCoord
}

func (f Generator) genRoutine(wg *sync.WaitGroup, rno int) {

	// Keeping as many recalculated values outside of the for loops as possible.
	offsets := make([]float64, f.samples)
	for sample := 0; sample < f.samples; sample++ {
		offsets[sample] = (1 + float64(2*sample) - float64(f.samples)) / float64(2*(f.samples))
	}
	samplesSquared := float64(f.samples * f.samples)
	routines := f.routines
	size := f.width * f.height

	for i := rno; i < size; i = i + routines {
		xPix := i % f.width
		yPix := i / f.width

		R, G, B, A := 0.0, 0.0, 0.0, 0.0

		for xSample := 0; xSample < f.samples; xSample++ {
			for ySample := 0; ySample < f.samples; ySample++ {
				xCoord, yCoord := f.pixToCoord(float64(xPix)+offsets[xSample], float64(yPix)+offsets[ySample])

				r, g, b, a := f.fn(xCoord, yCoord, f.iterationCap)

				R, G, B, A = R+r, G+g, B+b, A+a
			}
		}

		f.Img.Set(xPix, yPix,
			color.RGBA{
				uint8(R / samplesSquared),
				uint8(G / samplesSquared),
				uint8(B / samplesSquared),
				uint8(A / samplesSquared)})
	}

	wg.Done()
}
