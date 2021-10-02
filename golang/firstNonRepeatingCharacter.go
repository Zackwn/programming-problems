package main

func FirstNonRepeating(str string) string {
	chars := make(map[rune]int, len(str))
	for _, c := range str {
		chars[lower(c)]++
	}
	for _, c := range str {
		if chars[lower(c)] == 1 {
			return string(c)
		}
	}
	return ""
}

func lower(c rune) rune {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}
