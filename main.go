package main

import (
	"fmt"
)

func solution(sequence []int, k int) []int {
	n := len(sequence)
	var ans []int
	i, j := 0, 1
	sum := sequence[i] + sequence[j]
	minLen := n + 1
	for i < n {
		if sum == k {
			if j-i+1 < minLen {
				ans = []int{i, j}
				minLen = j - i + 1
			}
			j++
			if j == n {
				break
			}
			sum += sequence[j]
		} else if sum < k {
			j++
			if j == n {
				break
			}
			sum += sequence[j]
		} else {
			sum -= sequence[i]
			i++
		}
	}
	return ans
}

func main() {
	sequence := []int{1, 2, 3, 4, 5, 7, 9, 10, 13, 15, 20}
	k := 23
	ans := solution(sequence, k)
	fmt.Println(ans)
}
