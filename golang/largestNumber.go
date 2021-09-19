package main

func largestNumber(n int) int {
	o := 0
	for n > 0 {
		o = (o * 10) + 9
		n--
	}
	return o
}
