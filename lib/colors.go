package lib

import (
	"math"
)

type colorFunc func(iterations, iterationCap int, z, c complex) (R, G, B, A float64)

type colorScheme struct {
	Name string
	Fn   colorFunc
}

var defaultColors = []colorScheme{
	simpleGrayscale,
	simpleGrayscale,
	wackyGrayscale,
	wackyRainbow,
	zGrayscale,
	smoothGrayscale,
	smoothColor,
	smoothColor2,
}

var simpleGrayscale = colorScheme{
	Name: "simpleGrayscale",
	Fn: func(iterations, iterationCap int, z, c complex) (R, G, B, A float64) {
		col := float64(255*iterations) / float64(iterationCap)
		return col, col, col, 255
	},
}

// returns a color func that cycles through the set of colors passed in
func wacky(colors [][4]float64) colorFunc {
	return func(iterations, iterationCap int, z, c complex) (R, G, B, A float64) {
		key := iterations % len(colors)
		color := colors[key]
		return color[0], color[1], color[2], color[3]
	}
}

var wackyGrayscale = colorScheme{
	Name: "wackyGrayscale",
	Fn: wacky([][4]float64{
		[4]float64{0, 0, 0, 255},
		[4]float64{255, 255, 255, 255},
	}),
}

var wackyRainbow = colorScheme{
	Name: "wackyRainbow",
	Fn: wacky([][4]float64{
		[4]float64{84, 110, 98, 255},   // grey-green
		[4]float64{79, 127, 135, 255},  // turq
		[4]float64{110, 93, 158, 255},  // purp
		[4]float64{167, 125, 197, 255}, // pale purp
		[4]float64{255, 142, 145, 255}, // coral
		[4]float64{233, 186, 90, 255},  // orange
		[4]float64{231, 236, 128, 255}, // pale yellow
		[4]float64{135, 175, 95, 255},  // neon green
	}),
}

var zGrayscale = colorScheme{
	Name: "zGrayscale",
	Fn: func(iterations, iterationCap int, z, c complex) (R, G, B, A float64) {
		col := 255.0 * (math.Mod(z.abs(), 2.0) / 2.0)
		return col, col, col, 255
	},
}

var smoothGrayscale = colorScheme{
	Name: "smoothGrayscale",
	Fn: func(iterations, iterationCap int, z, c complex) (R, G, B, A float64) {
		z = z.mul(z).add(c)
		iterations++
		z = z.mul(z).add(c)
		iterations++

		i := float64(iterations)

		if iterations < iterationCap {
			i = i - (math.Log(math.Log(z.abs())) / math.Log(2))
		}

		if int(math.Floor(i))%2 == 0 {
			col := 255 * (math.Mod(i, 1))
			return col, col, col, 255
		}
		col := 255 - (255 * math.Mod(i, 1))
		return col, col, col, 255

	},
}

var smoothColor = colorScheme{
	Name: "smoothColor",
	Fn: func(iterations, iterationCap int, z, c complex) (R, G, B, A float64) {
		z = z.mul(z).add(c)
		iterations++
		z = z.mul(z).add(c)
		iterations++

		i := float64(iterations)

		if iterations < iterationCap {
			i = i - (math.Log(math.Log(z.abs())) / math.Log(2))
		}

		nu := math.Mod(i, 1)

		switch {
		case int(math.Floor(i))%3 == 0:
			return 255 * nu, 255 * (1 - nu), 255, 255
		case int(math.Floor(i))%3 == 1:
			return 255, 255 * nu, 255 * (1 - nu), 255
		case int(math.Floor(i))%3 == 2:
			return 255 * (1 - nu), 255, 255 * nu, 255
		}
		return 0, 0, 0, 255
	},
}

var smoothColor2 = colorScheme{
	Name: "smoothColor2",
	Fn: func(iterations, iterationCap int, z, c complex) (R, G, B, A float64) {
		z = z.mul(z).add(c)
		iterations++
		z = z.mul(z).add(c)
		iterations++

		i := float64(iterations)

		if iterations < iterationCap {
			i = i - (math.Log(math.Log(z.abs())) / math.Log(2))
		}

		nu := math.Mod(i, 1)

		switch {
		case int(math.Floor(i))%3 == 0:
			return 255 * (1 - nu), 255 * nu, 0, 255
		case int(math.Floor(i))%3 == 1:
			return 0, 255 * (1 - nu), 255 * nu, 255
		case int(math.Floor(i))%3 == 2:
			return 255 * nu, 0, 255 * (1 - nu), 255
		}
		return 0, 0, 0, 255
	},
}
