package main

import (
	"fmt"
	"math/big"
	"os"
)

var fractals = map[string]interface{}{
	"mandelbrot": map[string]interface{}{
		"description": "Classic mandelbrot function.",
		"constants":   0,
		"func":        mandelbrot,
		"colourfuncs": map[string]interface{}{
			"default":         simpleGreyscale,
			"simplegreyscale": simpleGreyscale,
			"whackygreyscale": whackyGrayscale,
			"zgreyscale":      zGreyscale,
			"smoothgreyscale": smoothGreyscale,
			"smoothcolour":    smoothColour,
			"smoothcolour2":   smoothColour2,
		},
	},
	"julia": map[string]interface{}{
		"description": "Classic Julia function.\nThe two constants are the real and imaginary components of C.",
		"constants":   2,
		"func":        julia,
		"colourfuncs": map[string]interface{}{
			"default":         simpleGreyscale,
			"simplegreyscale": simpleGreyscale,
			"whackygreyscale": whackyGrayscale,
			"zgreyscale":      zGreyscale,
			"smoothgreyscale": smoothGreyscale,
			"smoothcolour":    smoothColour,
			"smoothcolour2":   smoothColour2,
		},
	},
	"burningship": map[string]interface{}{
		"description": "Classic burning ship function.",
		"constants":   0,
		"func":        burningShip,
		"colourfuncs": map[string]interface{}{
			"default":         simpleGreyscaleShip,
			"simplegreyscale": simpleGreyscaleShip,
			"whackygreyscale": whackyGrayscaleShip,
		},
	},
}

func getFractalFunction(fractalName, colouringFuncName string, constants []float64) func(*big.Float, *big.Float, int) (R, G, B, A float64) {
	fractalFuncUnasserted, valid := fractals[fractalName] //Asserted after validation because if the fractal function's wrong, we'd try to assert nil.
	if valid != true {
		fmt.Println("Invalid fractal function.")
		os.Exit(1)
	}
	fractalFunc := fractalFuncUnasserted.(map[string]interface{})["func"].(func(interface{}, []float64) func(*big.Float, *big.Float, int) (float64, float64, float64, float64))

	if len(constants) != fractals[fractalName].(map[string]interface{})["constants"].(int) {
		fmt.Println("Invalid amount of constants.")
		os.Exit(1)
	}

	colouringFunc, valid := fractals[fractalName].(map[string]interface{})["colourfuncs"].(map[string]interface{})[colouringFuncName]

	if valid != true {
		fmt.Println("Invalid colouring function.")
		os.Exit(1)
	}

	return fractalFunc(colouringFunc, constants)
}

func mandelbrot(colourFuncUnasserted interface{}, constants []float64) func(*big.Float, *big.Float, int) (R, G, B, A float64) {
	colourFunc := colourFuncUnasserted.(func(int, int, complex, complex) (R, G, B, A float64))

	getMandelPoint := func(xCoord, yCoord *big.Float, iterationCap int) (R, G, B, A float64) {
		c := complex{xCoord, yCoord}
		z := complex{new(big.Float), new(big.Float)}
		iterations := 0

		for iterations = 0; z.abs().Cmp(big.NewFloat(2.0)) <= 0 && iterations < iterationCap; iterations++ {
			z = z.mul(z).add(c)
		}

		return colourFunc(iterations, iterationCap, z, c)
	}
	return getMandelPoint
}

func julia(colourFuncUnasserted interface{}, constants []float64) func(*big.Float, *big.Float, int) (R, G, B, A float64) {
	colourFunc := colourFuncUnasserted.(func(int, int, complex, complex) (R, G, B, A float64))

	getJuliaPoint := func(xCoord, yCoord *big.Float, iterationCap int) (R, G, B, A float64) {
		c := complex{big.NewFloat(constants[0]), big.NewFloat(constants[1])}
		z := complex{xCoord, yCoord}
		iterations := 0

		for iterations = 0; z.mul(z).add(c).abs().Cmp(big.NewFloat(2.0)) <= 0 && iterations < iterationCap; iterations++ {
			z = z.mul(z).add(c)
		}

		return colourFunc(iterations, iterationCap, z, c)
	}
	return getJuliaPoint
}

func burningShip(colourFuncUnasserted interface{}, constants []float64) func(*big.Float, *big.Float, int) (R, G, B, A float64) {
	colourFunc := colourFuncUnasserted.(func(int, int, complex) (R, G, B, A float64))

	getShipPoint := func(xCoord, yCoord *big.Float, iterationCap int) (R, G, B, A float64) {
		z := complex{big.NewFloat(0), big.NewFloat(0)}
		iterations := 0

		for iterations = 0; z.abs().Cmp(big.NewFloat(10.0)) <= 0 && iterations < iterationCap; iterations++ {
			zrzr := new(big.Float).Mul(z.real, z.real)
			zizi := new(big.Float).Mul(z.imag, z.imag)
			newReal := new(big.Float).Sub(zrzr, zizi)
			newReal.Add(newReal, xCoord)

			newImag := new(big.Float).Mul(new(big.Float).Abs(z.real), new(big.Float).Abs(z.imag))
			newImag.Mul(newImag, big.NewFloat(2.0))
			newImag.Add(newImag, yCoord)

			z.imag = newImag
			z.real = newReal
		}

		return colourFunc(iterations, iterationCap, z)
	}
	return getShipPoint
}
