package main

import (
	"fmt"
	"math"
)

const (
	alpha = 0.05
	power = 0.8
	z     = 1.96
)

func main() {
	p1 := 0.05
	n := 10000.0

	p2 := 0.1
	m := 100.0

	x := p1 - p2

	delta := math.Sqrt(p1*(1-p1)/n + p2*(1-p2)/m)
	fmt.Println(x / delta)
}
