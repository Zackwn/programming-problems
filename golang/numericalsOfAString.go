package main

func Numericals(s string) string {
	occur := make(map[rune]int)
	result := ""
	for _, c := range s {
		occur[c]++
		result += str(occur[c])
	}
	return result
}

func str(n int) string {
	s := ""
	for n > 0 {
		s = string(rune(n%10+'0')) + s
		n /= 10
	}
	return s
}
