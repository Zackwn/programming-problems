package main

func maxArea(height []int) int {
	max := 0
	l, r := 0, len(height)-1
	for l < r {
		area := min(height[l], height[r]) * (r - l)
		if area > max {
			max = area
		}
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return max
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
