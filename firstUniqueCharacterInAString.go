package main

func firstUniqChar(s string) int {
	chars := [26]int{}
	for _, char := range s {
		chars[char-'a']++
	}
	for i, char := range s {
		if chars[char-'a'] == 1 {
			return i
		}
	}
	return -1
}
