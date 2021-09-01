package main

func areEquallyStrong(yl int, yr int, fl int, fr int) bool {
	return max(yl, yr) == max(fl, fr) && min(yl, yr) == min(fl, fr)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
