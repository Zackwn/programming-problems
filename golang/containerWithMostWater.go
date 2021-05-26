package main

func maxArea(height []int) int {
	water := 0
	l, r := 0, len(height)-1
	for l < r {
		if min(height[l], height[r])*(r-l) > water {
			water = min(height[l], height[r]) * (r - l)
		}
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return water
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
