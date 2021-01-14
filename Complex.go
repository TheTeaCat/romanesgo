package main

import (
	"math/big"
)

type complex struct {
	real *big.Float
	imag *big.Float
}

func (c complex) abs() *big.Float {
	return new(big.Float).Sqrt(
		new(big.Float).Add(
			new(big.Float).Mul(c.real, c.real),
			new(big.Float).Mul(c.imag, c.imag)))
}

func (c complex) mul(d complex) (e complex) {
	e.real = new(big.Float).Sub(
		new(big.Float).Mul(c.real, d.real),
		new(big.Float).Mul(c.imag, d.imag))
	e.imag = new(big.Float).Add(
		new(big.Float).Mul(c.real, d.imag),
		new(big.Float).Mul(c.imag, d.real))
	return e
}

func (c complex) div(d complex) (e complex) {
	//Calculate the denom for both real and imag parts first...
	denom := new(big.Float).Add(
		new(big.Float).Mul(d.real, d.real),
		new(big.Float).Mul(d.imag, d.imag))

	//Real numerator (cr*dr + ci*di)
	e.real.Quo(
		new(big.Float).Add(
			new(big.Float).Mul(c.real, d.real),
			new(big.Float).Mul(c.imag, d.imag)),
		denom)

	//Imag numerator (ci*dr - cr*di)
	e.imag.Quo(
		new(big.Float).Sub(
			new(big.Float).Mul(c.imag, d.real),
			new(big.Float).Mul(c.real, d.imag)),
		denom)

	return e
}

func (c complex) add(d complex) (e complex) {
	e.real = new(big.Float).Add(c.real, d.real)
	e.imag = new(big.Float).Add(c.imag, d.imag)
	return e
}
