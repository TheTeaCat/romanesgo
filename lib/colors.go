package lib

import "math"

type colorFunc func(iterations, iterationCap int, z, c complex) (R, G, B, A float64)

type colorScheme struct {
	Name string
	Fn   colorFunc
}

var defaultColors = []colorScheme{
	simpleGreyscale,
	simpleGreyscale,
	wackyGrayscale,
	zGreyscale,
	smoothGreyscale,
	smoothColor,
	smoothColor2,
}

var simpleGreyscale = colorScheme{
	Name: "simpleGreyscale",
	Fn: func(iterations, iterationCap int, z, c complex) (R, G, B, A float64) {
		col := float64(255*iterations) / float64(iterationCap)
		return col, col, col, 255
	},
}

var simpleGreyscaleShip = colorScheme{
	Name: "simpleGreyscaleShip",
	Fn: func(iterations, iterationCap int, z, c complex) (R, G, B, A float64) {
		col := float64(255*iterations) / float64(iterationCap)
		return col, col, col, 255
	},
}

var wackyGrayscale = colorScheme{
	Name: "wackyGrayscale",
	Fn: func(iterations, iterationCap int, z, c complex) (R, G, B, A float64) {
		if iterations%2 == 0 {
			return 0, 0, 0, 255
		}
		return 255, 255, 255, 255
	},
}

var wackyGrayscaleShip = colorScheme{
	Name: "wackyGrayscaleShip",
	Fn: func(iterations, iterationCap int, z, c complex) (R, G, B, A float64) {
		if iterations%2 == 0 {
			return 0, 0, 0, 255
		}
		return 255, 255, 255, 255
	},
}

var zGreyscale = colorScheme{
	Name: "zGreyscale",
	Fn: func(iterations, iterationCap int, z, c complex) (R, G, B, A float64) {
		col := 255.0 * (math.Mod(z.abs(), 2.0) / 2.0)
		return col, col, col, 255
	},
}

var smoothGreyscale = colorScheme{
	Name: "smoothGreyscale",
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
