package main

// Input array is sorted
func twoSum(numbers []int, target int) []int {
	low, high := 0, len(numbers)-1
	for low < high {
		sum := numbers[low] + numbers[high]
		if sum == target {
			return []int{low + 1, high + 1}
		}

		if sum > target {
			high--
		} else {
			low++
		}
	}
	return []int{-1, -1}
}
