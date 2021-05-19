package main

func firstUniqChar(s string) int {
	chars := make(map[byte]uint, 26)
	for i := 0; i < len(s); i++ {
		chars[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		v := chars[s[i]]
		if v == 1 {
			return i
		}
	}
	return -1
}
