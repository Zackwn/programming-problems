package main

func CountBits(n uint) (count int) {
	for n > 0 {
		if n%2 == 1 {
			count++
		}
		n = n / 2
	}
	return
}
