package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println(fourSum(nums, target))
	nums = []int{2, 2, 2, 2, 2}
	target = 8
	fmt.Println(fourSum(nums, target))
	nums = []int{-2, -1, -1, 1, 1, 2, 2}
	target = 0
	fmt.Println(fourSum(nums, target))
	nums = []int{-3, -1, 0, 2, 4, 5}
	target = 2
	fmt.Println(fourSum(nums, target))
	nums = []int{-1, 0, 1, 2, -1, -4}
	target = -1
	fmt.Println(fourSum(nums, target))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	solutions := [][]int{}
	for i := 0; i < len(nums)-3; i++ {
		if i == 0 || (nums[i] != nums[i-1]) {
			for j := i + 1; j < len(nums)-2; j++ {
				if j == i+1 || (nums[j] != nums[j-1]) {
					low, high := j+1, len(nums)-1
					for low < high {
						sum := nums[i] + nums[j] + nums[low] + nums[high]
						if sum == target {
							solutions = append(solutions, []int{nums[i], nums[j], nums[low], nums[high]})
							for low < high && nums[low] == nums[low+1] {
								low++
							}
							for high > low && nums[high] == nums[high-1] {
								high--
							}
						}
						if sum > target {
							high--
						} else {
							low++
						}
					}
				}
			}
		}
	}
	return solutions
}

// failing tests
func fourSum2(nums []int, target int) [][]int {
	sort.Ints(nums)
	// fmt.Println(nums)
	solutions := [][]int{}

	for i := 0; i < len(nums); i++ {
		if i == 0 || (nums[i] != nums[i-1]) {

			n1 := nums[i]
			need := target - n1

			low, high := i+1, len(nums)-1
			for low+1 < high {
				fill := need - (nums[low] + nums[high])
				// fmt.Println(n1, need, fill, low, high)
				// if fill <= nums[high] || fill >= nums[low] {
				j := low + 1
				for j < high {
					if nums[j] == fill {
						solutions = append(solutions, []int{n1, nums[low], nums[j], nums[high]})
						for high > low && nums[high] == nums[high-1] {
							high--
						}
						for low < high && nums[low] == nums[low+1] {
							low++
						}
						break
					}
					j++
				}
				//j == high && n1+need+fill >= target
				// nums[low]+nums[high] > need
				if j == high && n1+need+fill <= target {
					high--
				} else {
					low++
				}
			}
		}
	}

	return solutions
}
