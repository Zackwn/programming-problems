package main

import (
	"fmt"
)

func main() {
	fmt.Println(IntersectionOfNumbers([]int{2, 4, 4, 2}, []int{2, 4}))
	fmt.Println(IntersectionOfNumbers([]int{1, 2, 3, 3}, []int{3, 3}))
	fmt.Println(IntersectionOfNumbers([]int{2, 4, 6, 8}, []int{1, 3, 5, 7}))
}

func IntersectionOfNumbers(a, b []int) []int {
	result := []int{}
	visitedNums := map[int]bool{}
	for _, aN := range a {
		_, visited := visitedNums[aN]
		if visited {
			continue
		}
		for _, bN := range b {
			if aN == bN {
				visitedNums[aN] = true
				result = append(result, aN)
				break
			}
		}
	}
	return result
}
