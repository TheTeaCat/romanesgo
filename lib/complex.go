package lib

import (
	"math"
)

type complex struct {
	real float64
	imag float64
}

func (c complex) abs() float64 {
	return math.Sqrt((c.real * c.real) + (c.imag * c.imag))
}

func (c complex) add(d complex) (e complex) {
	e.real = c.real + d.real
	e.imag = c.imag + d.imag
	return e
}

func (c complex) sub(d complex) (e complex) {
	e.real = c.real - d.real
	e.imag = c.imag - d.imag
	return e
}

func (c complex) div(d complex) (e complex) {
	e.real = ((c.real * d.real) + (c.imag * d.imag)) / ((d.real * d.real) + (d.imag * d.imag))
	e.imag = ((c.imag * d.real) - (c.real * d.imag)) / ((d.real * d.real) + (d.imag * d.imag))
	return e
}

func (c complex) mul(d complex) (e complex) {
	e.real = (c.real * d.real) - (c.imag * d.imag)
	e.imag = (c.real * d.imag) + (c.imag * d.real)
	return e
}

//Lazy/shorthand helper. This is faster than pow(2).
func (c complex) sq() complex {
	return c.mul(c)
}

func (c complex) pow(n float64) (e complex) {
	rN := math.Pow(c.real*c.real+c.imag*c.imag, n/2)
	nTheta := n * math.Atan2(c.imag, c.real)
	e.real = rN * math.Cos(nTheta)
	e.imag = rN * math.Sin(nTheta)
	return e
}

func (c complex) sin() (e complex) {
	e.real = math.Sin(c.real) * math.Cosh(c.imag)
	e.imag = math.Cos(c.real) * math.Sinh(c.imag)
	return e
}

func (c complex) cos() (e complex) {
	e.real = math.Cos(c.real) * math.Cosh(c.imag)
	e.imag = math.Sin(c.real) * math.Sinh(c.imag)
	return e
}

func (c complex) exp() (e complex) {
	e.real = math.Exp(c.real) * math.Cos(c.imag)
	e.imag = math.Exp(c.real) * math.Sin(c.imag)
	return e
}

func (c complex) conj() (e complex) {
	e.real = c.real
	e.imag = -c.imag
	return e
}
