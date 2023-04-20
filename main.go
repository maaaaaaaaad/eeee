package main

import (
	"fmt"
	"math"
)

func main() {
	r := 5.0
	h := 10.0
	n := 1000

	l := math.Sqrt(r*r + h*h)

	f := func(x float64) float64 {
		rx := (r * (h - x)) / h
		ax := math.Pi * rx * rx
		vx := (1.0 / 3.0) * ax * (h - x)
		return vx
	}

	a := 0.0
	b := h
	dx := (b - a) / float64(n)
	sum := (f(a) + f(b)) / 2.0
	for i := 1; i < n; i++ {
		x := a + float64(i)*dx
		sum += f(x)
	}
	integral := dx * sum

	volume := integral * l / 3.0

	fmt.Printf("The volume of the cone is %.2f\n", volume)
}
