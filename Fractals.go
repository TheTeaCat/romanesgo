package main

import (
	"fmt"
	"math"
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

func getFractalFunction(fractalName, colouringFuncName string, constants []float64) func(float64, float64, int) (R, G, B, A float64) {
	fractalFuncUnasserted, valid := fractals[fractalName] //Asserted after validation because if the fractal function's wrong, we'd try to assert nil.
	if valid != true {
		fmt.Println("Invalid fractal function.")
		os.Exit(1)
	}
	fractalFunc := fractalFuncUnasserted.(map[string]interface{})["func"].(func(interface{}, []float64) func(float64, float64, int) (float64, float64, float64, float64))

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

func mandelbrot(colourFuncUnasserted interface{}, constants []float64) func(float64, float64, int) (R, G, B, A float64) {
	colourFunc := colourFuncUnasserted.(func(int, int, complex, complex) (R, G, B, A float64))

	getMandelPoint := func(xCoord, yCoord float64, iterationCap int) (R, G, B, A float64) {
		c := complex{xCoord, yCoord}
		z := complex{0.0, 0.0}
		iterations := 0

		for iterations = 0; z.abs() <= 2 && iterations < iterationCap; iterations++ {
			z = z.mul(z).add(c)
		}

		return colourFunc(iterations, iterationCap, z, c)
	}
	return getMandelPoint
}

func julia(colourFuncUnasserted interface{}, constants []float64) func(float64, float64, int) (R, G, B, A float64) {
	colourFunc := colourFuncUnasserted.(func(int, int, complex, complex) (R, G, B, A float64))

	getJuliaPoint := func(xCoord, yCoord float64, iterationCap int) (R, G, B, A float64) {
		c := complex{constants[0], constants[1]}
		z := complex{xCoord, yCoord}
		iterations := 0

		for iterations = 0; z.mul(z).add(c).abs() <= 2 && iterations < iterationCap; iterations++ {
			z = z.mul(z).add(c)
		}

		return colourFunc(iterations, iterationCap, z, c)
	}
	return getJuliaPoint
}

func burningShip(colourFuncUnasserted interface{}, constants []float64) func(float64, float64, int) (R, G, B, A float64) {
	colourFunc := colourFuncUnasserted.(func(int, int, complex) (R, G, B, A float64))

	getShipPoint := func(xCoord, yCoord float64, iterationCap int) (R, G, B, A float64) {
		z := complex{0, 0}
		iterations := 0

		for iterations = 0; z.abs() <= 10 && iterations < iterationCap; iterations++ {
			newReal := (z.real * z.real) - (z.imag * z.imag) + xCoord
			z.imag = (2 * math.Abs(z.real) * math.Abs(z.imag)) + yCoord
			z.real = newReal
		}

		return colourFunc(iterations, iterationCap, z)
	}
	return getShipPoint
}
