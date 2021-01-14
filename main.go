package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/karlkeefer/romanesgo/lib"
)

func main() {
	fractalName := flag.String("ff", "none", "fractal")
	var constants flagConstants
	flag.Var(&constants, "c", "constants")
	iterations := flag.Int("i", 128, "maximum iterations")
	colorName := flag.String("cf", "default", "coloring function")
	xCentre := flag.Float64("x", 0, "central x coord")
	yCentre := flag.Float64("y", 0, "central y coord")
	zoom := flag.Float64("z", 1, "zoom factor")
	width := flag.Int("w", 1000, "image width")
	height := flag.Int("h", 1000, "image height")
	samples := flag.Int("ss", 1, "supersampling factor")
	routines := flag.Int("r", runtime.NumCPU(), "goroutines used")
	fn := flag.String("fn", "temp.png", "filename")
	flag.Parse()

	args := flag.Args()

	if len(args) > 0 && args[0] == "help" {
		handleHelp(args)
	} else {
		pointFunc, err := lib.GetPointFunc(*fractalName, *colorName, constants)
		fatal(err)

		fmt.Print("\n\tFractal (ff):\t\t", *fractalName,
			"\n\tConstants (c):\t\t", constants.String(),
			"\n\tMax Iterations (i):\t", *iterations,
			"\n\tColoring function (cf):\t", *colorName,
			"\n\tCentre x Coord (x):\t", *xCentre,
			"\n\tCentre y Coord (y):\t", *yCentre,
			"\n\tZoom factor (z):\t", *zoom,
			"\n\tImage Width (w):\t", *width,
			"\n\tImage Height (h):\t", *height,
			"\n\tSupersampling (ss):\t", *samples,
			"\n\tRoutines (r):\t\t", *routines,
			"\n\tFilename (png) (fn):\t", *fn, "\n\n")

		gen := lib.NewGenerator(*width, *height, *routines, *iterations, *samples, *xCentre, -*yCentre, *zoom, pointFunc)

		newFile, err := os.Create(*fn)
		fatal(err)

		timeIt(func() {
			gen.Generate()

			err = png.Encode(newFile, gen.Img)
			fatal(err)
		})
	}
}

func handleHelp(args []string) {
	if len(args) == 1 {
		fmt.Print(`Do "romanesgo help {Fractal Name}" for further info on a particular fractal function.`)
		fmt.Println("Fractals:")
		for _, f := range lib.Fractals {
			fmt.Println("\t", f.Name)
		}

		fmt.Println("\nFlags:")
		flag.PrintDefaults()
	} else if len(args) == 2 {
		frac, err := lib.GetFractal(args[1])
		fatal(err)
		fmt.Println(frac)
	}
}

func fatal(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}
}

func timeIt(fn func()) {
	start := time.Now()
	fn()
	fmt.Println("Done in", time.Since(start))
}

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
