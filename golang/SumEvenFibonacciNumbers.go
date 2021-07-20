package main

func SumEvenFibonacci(limit int) int {
	l, r := 0, 1
	var c int
	sum := 0
	for c <= limit {
		c = l + r
		if c%2 == 0 {
			sum += c
		}
		l = r
		r = c
	}
	return sum
}
