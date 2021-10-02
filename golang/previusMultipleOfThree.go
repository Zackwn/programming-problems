package main

func PrevMultOfThree(n int) interface{} {
	for n > 0 && n%3 != 0 {
		n /= 10
	}
	if n == 0 {
		return nil
	}
	return n
}
