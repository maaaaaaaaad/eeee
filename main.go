package main

import (
	"fmt"
)

func main() {
	a := []float64{1.0, 2.0, 3.0}
	b := []float64{4.0, 5.0, 6.0}

	dotProduct := 0.0
	for i := 0; i < len(a); i++ {
		dotProduct += a[i] * b[i]
	}

	fmt.Printf("두 벡터의 내적 %.2f", dotProduct)
}
