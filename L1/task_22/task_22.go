package main

import (
	"math/big"
)

func main() {
	a, _ := new(big.Int).SetString("999999999", 10)
	b, _ := new(big.Int).SetString("999999999999", 10)

	var sum, mult, quo, diff big.Int

	sum.Add(a, b)
	diff.Sub(a, b)
	mult.Mul(a, b)
	quo.Div(b, a)

	println("sum:", sum.String())
	println("diff:", diff.String())
	println("mult:", mult.String())
	println("quo:", quo.String())
}
