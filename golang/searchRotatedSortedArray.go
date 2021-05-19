package main

func search(nums []int, target int) int {
	indexStart := 0
	indexEnd := len(nums) - 1

	for indexStart <= indexEnd {
		mid := indexStart + (indexEnd-indexStart)/2

		if nums[mid] == target {
			return mid
		}

		if nums[indexStart] <= nums[mid] {
			if target >= nums[indexStart] && target <= nums[mid] {
				indexEnd = mid - 1
			} else {
				indexStart = mid + 1
			}
		} else {
			// nums[indexStart] > nums[mid]
			if target > nums[mid] && target <= nums[indexEnd] {
				indexStart = mid + 1
			} else {
				indexEnd = mid - 1
			}
		}
	}

	return -1
}
