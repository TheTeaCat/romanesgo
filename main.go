package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

type flagConstants []float64

func (f *flagConstants) String() string {
	str := ""
	for key, val := range *f {
		if key > 0 {
			str = str + ", "
		}
		str = str + strconv.FormatFloat(val, 'f', -1, 64)
	}
	return str
}

func (f *flagConstants) Set(value string) error {
	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}
	*f = append(*f, val)
	return nil
}

func sortMapKeys(m map[string]interface{}) []string {
	var sorted []string
	for key := range m {
		sorted = append(sorted, key)
	}
	sort.Strings(sorted)
	return sorted
}

func main() {
	fractalFunc := flag.String("ff", "none", "fractal")
	var constants flagConstants
	flag.Var(&constants, "c", "constants")
	iterations := flag.Int("i", 128, "maximum iterations")
	precision := flag.Uint("p", 53, "precision of floating point numbers")
	colourFunc := flag.String("cf", "default", "colouring function")
	xCentre := flag.Float64("x", 0, "central x coord")
	yCentre := flag.Float64("y", 0, "central y coord")
	zoom := flag.Float64("z", 1, "zoom factor")
	width := flag.Int("w", 1000, "image width")
	height := flag.Int("h", 1000, "image height")
	samples := flag.Int("ss", 1, "supersampling factor")
	routines := flag.Int("r", runtime.NumCPU(), "goroutines used")
	fn := flag.String("fn", "temp.png", "filename")
	flag.Parse()

	if len(flag.Args()) == 1 && flag.Args()[0] == "help" {
		fmt.Println("\nDo \"romanesgo help {Fractal Name}\" for further info on a particular fractal function.\n")
		fmt.Println("Fractals:")

		fractalNames := sortMapKeys(fractals)
		for _, fractalName := range fractalNames {
			fmt.Println("\t", fractalName)
		}

		fmt.Println("\nFlags:")
		flag.PrintDefaults()

	} else if len(flag.Args()) == 2 && flag.Args()[0] == "help" {
		fractalName := strings.ToLower(flag.Args()[1])
		if _, valid := fractals[fractalName]; valid {
			fmt.Println("\nDescription:\n\n"+fractals[fractalName].(map[string]interface{})["description"].(string), "\n\nInfo:")
			fmt.Println("\tConstants:", fractals[fractalName].(map[string]interface{})["constants"].(int))
			fmt.Println("\tColouring functions:")

			colourFuncNames := sortMapKeys(fractals[fractalName].(map[string]interface{})["colourfuncs"].(map[string]interface{}))
			for _, funcName := range colourFuncNames {
				if funcName != "default" {
					fmt.Println("\t\t", funcName)
				}
			}

		} else {
			fmt.Println("\nUnrecognised fractal function name.")
		}
	} else if *fractalFunc == "none" {
		fmt.Println("\nDo \"romanesgo help\" for more info.")
	} else {
		fmt.Println("\n\tFractal (ff):\t\t", *fractalFunc,
			"\n\tConstants (c):\t\t", constants.String(),
			"\n\tMax Iterations (i):\t", *iterations,
			"\n\tPrecision (p):\t\t", *precision,
			"\n\tColouring function (cf):", *colourFunc,
			"\n\tCentre x Coord (x):\t", *xCentre,
			"\n\tCentre y Coord (y):\t", *yCentre,
			"\n\tZoom factor (z):\t", *zoom,
			"\n\tImage Width (w):\t", *width,
			"\n\tImage Height (h):\t", *height,
			"\n\tSupersampling (ss):\t", *samples,
			"\n\tRoutines (r):\t\t", *routines,
			"\n\tFilename (png) (fn):\t", *fn, "\n")
		f := getNewFractGen(*width, *height, *routines, *iterations, *precision, *xCentre, -*yCentre, *zoom)

		startTime := time.Now()
		f.generate(getFractalFunction(strings.ToLower(*fractalFunc), strings.ToLower(*colourFunc), constants), *samples)
		duration := time.Since(startTime)

		fmt.Println("\nTime taken:", duration)
		newFile, _ := os.Create(*fn)
		png.Encode(newFile, f.fractImg)
	}
}
