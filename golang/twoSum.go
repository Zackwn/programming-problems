func twoSum(nums []int, target int) []int {
	println(nums, target)
	for index, value := range nums {
		complain := target - value
		for i := index + 1; i < len(nums); i++ {
			if nums[i] == complain {
				return []int{index, i}
			}
		}
	}
	return []int{0, 0}
}