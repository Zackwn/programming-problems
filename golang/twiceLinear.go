package main

func DblLinear(n int) int {
	//fmt.Println("  ")
	u := make([]int, 1, n)
	u[0] = 1
	x, y := 0, 0
	for i := 0; i < n; i++ {
		nextX, nextY := 2*u[x]+1, 3*u[y]+1
		//fmt.Println(u[x], u[y])
		if nextY >= nextX {
			//fmt.Println("x", nextX)
			u = append(u, nextX)
			x++
			if nextY == nextX {
				y++
			}
		} else {
			//fmt.Println("y",nextY)
			u = append(u, nextY)
			y++
		}
		//fmt.Println(" ")
	}
	return u[n]
}
