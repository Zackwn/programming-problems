package main

func Pyramid(n int) [][]int {
	p := make([][]int, n) // pyramid
	s := make([]int, 0)   // step
	for i := 0; i < n; i++ {
		s = append(s, 1)
		p[i] = s
	}
	return p
}
