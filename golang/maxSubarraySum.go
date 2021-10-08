package main

// Kadane's algorithm <https://en.wikipedia.org/wiki/Maximum_subarray_problem>
func MaximumSubarraySum(numbers []int) int {
	sum, maxSum := 0, 0
	for _, n := range numbers {
		if sum+n > 0 {
			sum += n
		} else {
			sum = 0
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}
