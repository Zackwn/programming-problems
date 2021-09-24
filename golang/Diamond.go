package main

func Diamond(n int) string {
	if n <= 0 || n%2 == 0 {
		return ""
	}
	result := ""
	space, star := n/2, 1
	for star < n {
		result += s(" ", space) + s("*", star) + s(" ", space) + "\n"
		space--
		star += 2
	}
	for star > 0 {
		result += s(" ", space) + s("*", star) + s(" ", space) + "\n"
		space++
		star -= 2
	}
	return result
}

func s(c string, n int) string {
	r := ""
	for n > 0 {
		r += c
		n--
	}
	return r
}
