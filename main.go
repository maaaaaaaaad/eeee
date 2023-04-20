package main

import (
	"fmt"
	"math"
)

func main() {
	r := 5.0
	n := 1000
	a := 0.0
	b := 2.0 * math.Pi * r
	dx := (b - a) / float64(n)

	f := func(x float64) float64 {
		ax := math.Pi * (r*r - x*x)
		return ax
	}

	sum := (f(a) + f(b)) / 2.0
	for i := 1; i < n; i++ {
		x := a + float64(i)*dx
		sum += f(x)
	}
	integral := dx * sum

	volume := 4.0 / 3.0 * math.Pi * math.Pow(r, 3.0)
	volume = volume * integral / (math.Pi * math.Pow(r, 2.0))

	fmt.Printf(" %.2f\n", volume)
}
