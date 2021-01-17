package lib

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// Named errors might be helpful for matching outside this lib
var (
	ErrInvalidFractal      = errors.New("invalid fractal name")
	ErrInvalidColor        = errors.New("invalid color scheme name")
	ErrColorNotImplemented = errors.New("color scheme name valid but not implemented")
)

type fractalFunc func(color colorFunc, constants []float64) PointFunc

// Fractal gontains everything you need to get a colorized point function for our generator
type Fractal struct {
	Description        string
	Constants          int
	ColorSchemes       []string
	DefaultColorScheme string
	Fn                 fractalFunc
}

// String outputs basic info for the help screen
func (f Fractal) String() string {
	return fmt.Sprintf("%s\nColor Schemes: %s", f.Description, strings.Join(f.ColorSchemes, ", "))
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

	// colorNames should always be lowercased
	colorName = strings.ToLower(colorName)
	/* if colorName is the empty string, or "default", then we use the default
	   coloring function of the fractal.
	*/
	if colorName == "" || strings.ToLower(colorName) == "default" {
		colorName = frac.DefaultColorScheme
	}
	// check colorName is valid for this fractal
	colorValid := false
	for _, cn := range frac.ColorSchemes {
		if cn == colorName {
			colorValid = true
			break
		}
	}
	if !colorValid {
		return nil, ErrInvalidColor
	}
	/* Get color function from colorSchemes:
	   Could still fail if a colorscheme is named in the fractal object that
	   does not exist in the colorSchemes map.
	*/
	colorFunc, colorFuncExists := colorSchemes[colorName]
	if !colorFuncExists {
		return nil, ErrColorNotImplemented
	}

	return frac.Fn(colorFunc, constants), nil
}

// GetFractal returns a fractal if the name is valid
func GetFractal(fractalName string) (*Fractal, error) {
	frac, validFractal := Fractals[strings.ToLower(fractalName)]

	if !validFractal {
		return nil, ErrInvalidFractal
	}

	return frac, nil
}

// Fractals is an array of the available fractals in this program
var Fractals = map[string]*Fractal{
	"mandelbrot": &Fractal{
		Description:        "Classic mandelbrot function.",
		Constants:          0,
		ColorSchemes:       []string{"simplegrayscale", "wackygrayscale", "wackyrainbow", "zgrayscale", "smoothgrayscale", "smoothcolor", "smoothcolor2"},
		DefaultColorScheme: "simplegrayscale",
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

	"julia": &Fractal{
		Description:        "Classic Julia function.\nThe two constants are the real and imaginary components of C.",
		Constants:          2,
		ColorSchemes:       []string{"simplegrayscale", "wackygrayscale", "wackyrainbow", "zgrayscale", "smoothgrayscale", "smoothcolor", "smoothcolor2"},
		DefaultColorScheme: "simplegrayscale",
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

	"burningship": &Fractal{
		Description:        "Classic burning ship function.",
		Constants:          0,
		ColorSchemes:       []string{"simplegrayscale", "wackygrayscale"},
		DefaultColorScheme: "simplegrayscale",
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

	"collatz": &Fractal{
		Description:        "The Collatz fractal.\nThe constant value is the absolute value after which the sequence will be assumed to have escaped.",
		Constants:          0,
		ColorSchemes:       []string{"simplegrayscale", "wackygrayscale", "wackyrainbow"},
		DefaultColorScheme: "simplegrayscale",
		Fn: func(color colorFunc, constants []float64) PointFunc {
			return func(xCoord, yCoord float64, iterationCap int) (R, G, B, A float64) {
				z := complex{xCoord, yCoord}
				iterations := 0

				for iterations = 0; z.abs() < math.MaxFloat64 && iterations < iterationCap; iterations++ {
					cossq := z.mul(complex{math.Pi / 2, 0}).cos().sq()
					sinsq := z.mul(complex{math.Pi / 2, 0}).sin().sq()
					z = cossq.mul(z.mul(complex{0.5, 0})).add(
						sinsq.mul(z.mul(complex{3.0, 0}).add(complex{1.0, 0})))
				}

				return color(iterations, iterationCap, z, complex{0, 0})
			}
		},
	},

	"tricorn": &Fractal{
		Description:        "Classic tricorn function.",
		Constants:          0,
		ColorSchemes:       []string{"simplegrayscale", "wackygrayscale", "wackyrainbow", "zgrayscale"},
		DefaultColorScheme: "simplegrayscale",
		Fn: func(color colorFunc, constants []float64) PointFunc {
			return func(xCoord, yCoord float64, iterationCap int) (R, G, B, A float64) {
				c := complex{xCoord, yCoord}
				z := complex{0.0, 0.0}
				iterations := 0

				for iterations = 0; z.abs() <= 2 && iterations < iterationCap; iterations++ {
					z = z.conj()
					z = z.mul(z).add(c)
				}

				return color(iterations, iterationCap, z, c)
			}
		},
	},
}
