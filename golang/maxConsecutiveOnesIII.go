package main

func longestOnes(nums []int, k int) int {
	max, start, end := 0, 0, 0
	for start < len(nums) {
		if end < len(nums) && (nums[end] == 0 && k > 0 || nums[end] == 1) {
			if nums[end] == 0 {
				k--
			}
			end++
		} else {
			if (end == len(nums) && k > 0) || k == 0 {
				if end-start > max {
					max = end - start
				}
				if nums[start] == 0 {
					k++
				}
				start++
			}
		}
	}
	return max
}
