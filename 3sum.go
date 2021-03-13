func threeSum(nums []int) [][]int {
	/*
		for each element search for 3sum == 0 with left and right pointers
	*/
	triplets := [][]int{}
	sort.Sort(sort.IntSlice(nums))
	for index := 0; index < len(nums)-2; index++ {
		// skip repeated elements in a row
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}
		l, r := index+1, len(nums)-1
		for l < r {
			total := nums[index] + nums[l] + nums[r]
			if total == 0 {
				triplets = append(triplets, []int{nums[index], nums[l], nums[r]})
				// move to next left and right skiping repeateds
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				// move to the next pointers
				l++
				r--
			} else if total < 0 {
				/*array is sorted ascending
				so as the 3sum is smaller than 0 need
				greater numbers: moving ahead the left pointer
				*/
				l++
			} else {
				// total > 0
				/*array is sorted ascending
				so as the 3sum is greater than 0 need
				smaller numbers: moving right pointer back
				*/
				r--
			}
		}
	}
	return triplets
}