package main

func Finance(n int) int {
	var x, r int
	for n >= 0 {
		for i := 0; i <= n; i++ {
			r += x + i
		}
		x += 2
		n--
	}
	return r
}
