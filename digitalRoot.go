package main

func DigitalRoot(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n = n / 10
	}
	if sum > 9 {
		return DigitalRoot(sum)
	}
	return sum
}
