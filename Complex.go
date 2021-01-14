package main

import (
	"math/big"
)

type complex struct {
	real *big.Float
	imag *big.Float
}

//Returns the square of the absolute value (because square rooting is costly)
// (Squared Euclidean Distance)
func (c complex) absq() *big.Float {
	v := new(big.Float).Mul(c.real, c.real)
	return v.Add(v, new(big.Float).Mul(c.imag, c.imag))
}

func (c complex) abs() *big.Float {
	v := new(big.Float).Mul(c.real, c.real)
	v.Add(v, new(big.Float).Mul(c.imag, c.imag))
	return v.Sqrt(v)
}

func (c complex) mul(d complex) (e complex) {
	//er = cr*dr - ci*di
	e.real = new(big.Float).Mul(c.real, d.real)
	e.real.Sub(e.real, new(big.Float).Mul(c.imag, d.imag))

	//ei = cr*di + ci*dr
	e.imag = new(big.Float).Mul(c.real, d.imag)
	e.imag.Add(e.imag, new(big.Float).Mul(c.imag, d.real))

	return e
}

func (c complex) div(d complex) (e complex) {
	//Calculate the denom for both real and imag parts first...
	denom := new(big.Float).Mul(d.real, d.real)
	denom.Add(denom, new(big.Float).Mul(d.imag, d.imag))

	//Real numerator (cr*dr + ci*di)
	e.real = new(big.Float).Mul(c.real, d.real)
	e.real.Add(e.real, new(big.Float).Mul(c.imag, d.imag))
	e.real.Quo(e.real, denom)

	//Imag numerator (ci*dr - cr*di)
	e.imag = new(big.Float).Mul(c.imag, d.real)
	e.imag.Sub(e.imag, new(big.Float).Mul(c.real, d.imag))
	e.imag.Quo(e.imag, denom)

	return e
}

func (c complex) add(d complex) (e complex) {
	//trivial, er=cr+dr, ei=ci+di
	e.real = new(big.Float).Add(c.real, d.real)
	e.imag = new(big.Float).Add(c.imag, d.imag)
	return e
}
