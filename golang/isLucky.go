package main

func isLucky(n int) bool {
	digits := make([]int, 0)
	for n > 0 {
		digits = append(digits, n%10)
		n = n / 10
	}
	fhalf, shalf := 0, 0
	for i := 0; i < len(digits)/2; i++ {
		fhalf += digits[i]
		shalf += digits[len(digits)-i-1]
	}
	return fhalf == shalf
}
