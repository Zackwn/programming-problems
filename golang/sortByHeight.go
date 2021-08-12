package main

func sortByHeight(a []int) []int {
	h := make([]int, 0)
	for _, v := range a {
		if v == -1 {
			continue
		}
		h = append(h, v)
	}
	left, right := 0, 1
	for right < len(h) {
		if h[left] > h[right] {
			// swap
			h[left], h[right] = h[right], h[left]
			// backtrack
			for left-1 >= 0 && h[left-1] > h[left] {
				// swap
				h[left-1], h[left] = h[left], h[left-1]
				left--
			}
			// left tail right
			left = right - 1
		}
		// move pointers ahead
		right++
		left++
	}
	hi := 0
	for i, v := range a {
		if v != -1 {
			a[i] = h[hi]
			hi++
		}
	}
	return a
}
