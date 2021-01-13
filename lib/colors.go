package lib

import "math"

func simpleGreyscale(iterations, iterationCap int, z, c complex) (R, G, B, A float64) {
	col := float64(255*iterations) / float64(iterationCap)
	return col, col, col, 255
}

func simpleGreyscaleShip(iterations, iterationCap int, z complex) (R, G, B, A float64) {
	col := float64(255*iterations) / float64(iterationCap)
	return col, col, col, 255
}

func whackyGrayscale(iterations, iterationCap int, z, c complex) (R, G, B, A float64) {
	if iterations%2 == 0 {
		return 0, 0, 0, 255
	}
	return 255, 255, 255, 255
}

func whackyGrayscaleShip(iterations, iterationCap int, z complex) (R, G, B, A float64) {
	if iterations%2 == 0 {
		return 0, 0, 0, 255
	}
	return 255, 255, 255, 255
}

func zGreyscale(iterations, iterationCap int, z, c complex) (R, G, B, A float64) {
	col := 255.0 * (math.Mod(z.abs(), 2.0) / 2.0)
	return col, col, col, 255
}

func smoothGreyscale(iterations, iterationCap int, z, c complex) (R, G, B, A float64) {
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

}

func smoothColour(iterations, iterationCap int, z, c complex) (R, G, B, A float64) {
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
}

func smoothColour2(iterations, iterationCap int, z, c complex) (R, G, B, A float64) {
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
}
