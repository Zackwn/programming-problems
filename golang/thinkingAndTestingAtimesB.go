package main

func TestIt(a, b int) int {
	return sumDigits(a) * sumDigits(b)
}

func sumDigits(n int) (sum int) {
	for n > 0 {
		sum += n % 10
		n = n / 10
	}
	return
}
