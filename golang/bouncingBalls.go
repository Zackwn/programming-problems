package main

func BouncingBall(h, bounce, window float64) int {
	if h < 0 || bounce < 0 || bounce >= 1 || window >= h {
		return -1
	}
	times := 0
	for h > window {
		if h > window {
			times++
		}
		h = h * bounce
		if h > window {
			times++
		}
	}
	return times
}
