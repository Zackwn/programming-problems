package main

func shapeArea(n int) int {
	if n == 1 {
		return 1
	} else {
		a := 1
		for i := 2; i <= n; i++ {
			a += i*4 - 4
		}
		return a
	}
}
