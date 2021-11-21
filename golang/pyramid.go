package main

func Pyramid(n int) [][]int {
	pyd := make([][]int, n)
	for i := 1; i <= n; i++ {
		step := make([]int, i)
		for j := 0; j < i; j++ {
			step[j] = 1
		}
		pyd[i-1] = step
	}
	return pyd
}
