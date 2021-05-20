package main

import "fmt"

func main() {
	// output: 24
	fmt.Println(getMaxSumOfFiveContiguousElements([]int{5, 7, 1, 4, 3, 6, 2, 9, 2}))
}

func getMaxSumOfFiveContiguousElements(arr []int) int {
	if len(arr) < 5 {
		return 0
	}
	s, e := 0, 5
	sum := arr[s] + arr[s+1] + arr[s+2] + arr[s+3] + arr[s+4]
	maxsum := sum
	for e < len(arr) {
		sum -= arr[s]
		sum += arr[e]
		if sum > maxsum {
			maxsum = sum
		}
		s++
		e++
	}
	return maxsum
}
