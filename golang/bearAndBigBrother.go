package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println(BearAndBigBrother(a, b))
}

func BearAndBigBrother(a, b int) int {
	years := 0
	for b >= a {
		a *= 3
		b *= 2
		years++
	}
	return years
}
