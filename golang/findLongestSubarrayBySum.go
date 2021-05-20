package main

// sliding window algorithm
func findLongestSubarrayBySumBetterSolution(s int, arr []int) []int {
	rstart, rend := -1, -1
	start, end := 0, 0
	sum := 0
	for end < len(arr) {
		sum += arr[end]
		for start < len(arr) && sum > s {
			sum -= arr[start]
			start++
		}
		if sum == s && start-end <= rstart-rend {
			rstart = start
			rend = end
		}
		end++
	}
	if rend < 0 {
		return []int{-1}
	}
	return []int{rstart + 1, rend + 1}
}

// "brute force"
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
