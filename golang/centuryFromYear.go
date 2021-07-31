package main

func centuryFromYear(year int) int {
	y := 0
	c := 0
	y += year % 10
	year = year / 10
	y = addStart(y, year%10)
	year = year / 10
	c += year % 10
	year = year / 10
	c = addStart(c, year)
	if y != 0 {
		c++
	}
	return c
}

func addStart(n, a int) int {
	x := n
	n = a * 10
	n += x
	return n
}
