package main

import "fmt"

func main() {
	girlCount := 3
	boyCount := 5

	girlPermutation := factorial(girlCount)

	boyCombination := combination(boyCount, girlCount-1)

	result := girlPermutation * boyCombination

	fmt.Printf("여자 3명이 누구도 이웃하지 않는 경우의 수는 %d입니다.\n", result)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func combination(n int, r int) int {
	if r == 0 || r == n {
		return 1
	}

	if n <= 5 && r <= 3 {
		switch {
		case n == 5 && r == 1:
			return 5
		case n == 5 && r == 2:
			return 10
		case n == 5 && r == 3:
			return 10
		}
	}

	result := 1
	for i := 1; i <= r; i++ {
		result *= n - i + 1
		result /= i
	}
	return result
}
