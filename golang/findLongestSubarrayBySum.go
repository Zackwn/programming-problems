package main

func findLongestSubarrayBySum(s int, arr []int) []int {
	result := []int{-1}
	for i := 0; i < len(arr); i++ {
		sum := arr[i]
		j := i + 1
		for j < len(arr) && (sum < s || arr[j] == 0) {
			sum += arr[j]
			j++
		}
		if sum == s {
			if result[0] == -1 {
				result = []int{i + 1, j}
			} else if result[0]-result[1] > (i+1)-j {
				result = []int{i + 1, j}
			}
		}
	}
	return result
}
