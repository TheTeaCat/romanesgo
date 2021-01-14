package lib

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// Named errors might be helpful for matching outside this lib
var (
	ErrInvalidColor   = errors.New("invalid color scheme name")
	ErrInvalidFractal = errors.New("invalid fractal name")
)

type fractalFunc func(color colorFunc, constants []float64) PointFunc

// Fractal gontains everything you need to get a colorized point function for our generator
type Fractal struct {
	Name         string
	Description  string
	Constants    int
	ColorSchemes []colorScheme
	Fn           fractalFunc
}

// String outputs basic info for the help screen
func (f Fractal) String() string {
	colorSchemes := make([]string, len(f.ColorSchemes))
	for i, cs := range f.ColorSchemes {
		colorSchemes[i] = cs.Name
	}
	return fmt.Sprintf("%s\n%s\nColor Schemes: %s", f.Name, f.Description, strings.Join(colorSchemes, ", "))
}

// GetPointFunc will check for valid fractalname and colorname
// returns a pointFunc if we're good to go
func GetPointFunc(fractalName, colorName string, constants []float64) (PointFunc, error) {

	frac, err := GetFractal(fractalName)
	if err != nil {
		return nil, err
	}

	// check constants
	if len(constants) != frac.Constants {
		return nil, errors.New("invalid number of constants")
	}

	// check colorScheme
	var color *colorScheme
	for _, cs := range frac.ColorSchemes {
		if strings.ToLower(cs.Name) == strings.ToLower(colorName) {
			color = &cs
			break
		}
	}

	if color == nil {
		return nil, ErrInvalidColor
	}

	return frac.Fn(color.Fn, constants), nil
}

// GetFractal returns a fractal if the name is valid
func GetFractal(fractalName string) (frac *Fractal, err error) {
	// check fractalName
	for _, f := range Fractals {
		if f.Name == strings.ToLower(fractalName) {
			frac = &f
			break
		}
	}

	if frac == nil {
		return nil, ErrInvalidFractal
	}

	return frac, nil
}

// Fractals is an array of the available fractals in this program
var Fractals = []Fractal{
	Fractal{
		Name:         "mandelbrot",
		Description:  "Classic mandelbrot function.",
		Constants:    0,
		ColorSchemes: defaultColors,
		Fn: func(color colorFunc, constants []float64) PointFunc {
			return func(xCoord, yCoord float64, iterationCap int) (R, G, B, A float64) {
				c := complex{xCoord, yCoord}
				z := complex{0.0, 0.0}
				iterations := 0

				for iterations = 0; z.abs() <= 2 && iterations < iterationCap; iterations++ {
					z = z.mul(z).add(c)
				}

				return color(iterations, iterationCap, z, c)
			}
		},
	},

	Fractal{
		Name:         "julia",
		Description:  "Classic Julia function.\nThe two constants are the real and imaginary components of C.",
		Constants:    2,
		ColorSchemes: defaultColors,
		Fn: func(color colorFunc, constants []float64) PointFunc {
			return func(xCoord, yCoord float64, iterationCap int) (R, G, B, A float64) {
				c := complex{constants[0], constants[1]}
				z := complex{xCoord, yCoord}
				iterations := 0

				for iterations = 0; z.mul(z).add(c).abs() <= 2 && iterations < iterationCap; iterations++ {
					z = z.mul(z).add(c)
				}

				return color(iterations, iterationCap, z, c)
			}
		},
	},

	Fractal{
		Name:        "burningship",
		Description: "Classic burning ship function.",
		Constants:   0,
		ColorSchemes: []colorScheme{
			simpleGreyscaleShip,
			simpleGreyscaleShip,
			wackyGrayscaleShip,
		},
		Fn: func(color colorFunc, constants []float64) PointFunc {
			return func(xCoord, yCoord float64, iterationCap int) (R, G, B, A float64) {
				z := complex{0, 0}
				iterations := 0

				for iterations = 0; z.abs() <= 10 && iterations < iterationCap; iterations++ {
					newReal := (z.real * z.real) - (z.imag * z.imag) + xCoord
					z.imag = (2 * math.Abs(z.real) * math.Abs(z.imag)) + yCoord
					z.real = newReal
				}

				return color(iterations, iterationCap, z, complex{0, 0})
			}
		},
	},
}
