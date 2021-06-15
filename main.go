package main

import "github.com/s0xzwasd/false-positive-inspections/p"

type BImpl struct {
	p.AImpl
}

func (b BImpl) u(c int) int { return c }

func main() {
	var b BImpl

	var bi1 p.Iminterface
	bi1 = b

	bi2 := p.Iminterface(b)

	_, _, _ = b, bi1, bi2
}
