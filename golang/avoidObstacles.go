package main

func avoidObstacles(arr []int) int {
	r := 1
	for i := 0; i < len(arr); i++ {
		if arr[i]%r == 0 {
			i = -1
			r++
		}
	}
	return r
}
