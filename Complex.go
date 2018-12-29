package main

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

func (c complex) exp() (e complex) {
	e.real = math.Exp(c.real) * math.Cos(c.imag)
	e.imag = math.Exp(c.real) * math.Sin(c.imag)
	return e
}

func (c complex) mul(d complex) (e complex) {
	e.real = (c.real * d.real) - (c.imag * d.imag)
	e.imag = (c.real * d.imag) + (c.imag * d.real)
	return e
}

func (c complex) div(d complex) (e complex) {
	e.real = ((c.real * d.real) + (c.imag * d.imag)) / ((d.real * d.real) + (d.imag * d.imag))
	e.imag = ((c.imag * d.real) - (c.real * d.imag)) / ((d.real * d.real) + (d.imag * d.imag))
	return e
}

func (c complex) add(d complex) (e complex) {
	e.real = c.real + d.real
	e.imag = c.imag + d.imag
	return e
}
