package main

import (
	"fmt"
	"image"
	"image/color"
	"math/big"
	"sync"
)

type fractGen struct {
	fractImg     *image.NRGBA
	xPos         float64
	yPos         float64
	zoom         float64
	scaler       float64
	width        int
	height       int
	routines     int
	iterationCap int
	precision    uint
}

func getNewFractGen(width, height, routines, iterationCap int, precision uint, xPos, yPos, zoom float64) fractGen {
	newFractGen := fractGen{
		width:        width,
		height:       height,
		routines:     routines,
		iterationCap: iterationCap,
		precision:    precision,
		xPos:         xPos,
		yPos:         yPos,
		zoom:         zoom,
		fractImg:     image.NewNRGBA(image.Rect(0, 0, width, height)),
	}

	//Picks the smaller of the two dimensions (width and height) and uses that length in pixels as the length of 2 divided by the zoom factor as the scale for both axis.
	if newFractGen.width < newFractGen.height {
		newFractGen.scaler = float64(newFractGen.width)
	} else {
		newFractGen.scaler = float64(newFractGen.height)
	}

	return newFractGen

}

func (f fractGen) pixToCoord(xPix, yPix float64) (xCoord, yCoord *big.Float) {
	//Shifting the origin to the center of the image
	xCoord = new(big.Float).SetPrec(f.precision).Sub(big.NewFloat(xPix), big.NewFloat(float64(f.width)/2))
	yCoord = new(big.Float).SetPrec(f.precision).Sub(big.NewFloat(yPix), big.NewFloat(float64(f.height)/2))

	//Scaling the pixels down to values on the complex plane
	scaler := new(big.Float).Quo(big.NewFloat(2/f.scaler), big.NewFloat(f.zoom))
	xCoord.Mul(xCoord, scaler)
	yCoord.Mul(yCoord, scaler)

	//Shifting the origin over to the xPos and yPos on the complex plane
	xCoord.Add(xCoord, big.NewFloat(f.xPos))
	yCoord.Add(yCoord, big.NewFloat(f.yPos))

	return xCoord, yCoord
}

func (f fractGen) generate(pointFunc func(*big.Float, *big.Float, int) (R, G, B, A float64), samples int) {
	var wg sync.WaitGroup
	wg.Add(f.routines)

	for routine := 0; routine < f.routines; routine++ {
		go f.genRoutine(&wg, routine, samples, pointFunc)
	}

	wg.Wait()
}

func (f fractGen) genRoutine(wg *sync.WaitGroup, rno int, samples int, pointFunc func(*big.Float, *big.Float, int) (R, G, B, A float64)) {

	offsets := make([]float64, samples)

	for sample := 0; sample < samples; sample++ {
		offsets[sample] = (1 + float64(2*sample) - float64(samples)) / float64(2*(samples))
	}
	samplesSquared := float64(samples * samples) //Keeping as many recalculated values outside of the for loops as possible.

	routines := f.routines
	size := f.width * f.height
	for i := rno; i < size; i = i + routines {
		xPix := i % f.width
		yPix := i / f.width

		R, G, B, A := 0.0, 0.0, 0.0, 0.0

		for xSample := 0; xSample < samples; xSample++ {
			for ySample := 0; ySample < samples; ySample++ {
				xCoord, yCoord := f.pixToCoord(float64(xPix)+offsets[xSample], float64(yPix)+offsets[ySample])

				r, g, b, a := pointFunc(xCoord, yCoord, f.iterationCap)

				R, G, B, A = R+(r/samplesSquared),
					G+(g/samplesSquared),
					B+(b/samplesSquared),
					A+(a/samplesSquared)
			}
		}
		f.fractImg.Set(xPix, yPix, color.RGBA{uint8(R), uint8(G), uint8(B), uint8(A)})
	}
	fmt.Println("Routine", rno, "Done.")

	wg.Done()
}
