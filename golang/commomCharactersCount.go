package main

func commonCharacterCount(s1 string, s2 string) int {
	s1Chars := allChars(s1)
	s2Chars := allChars(s2)
	commomCharsCount := 0
	for c := range s1Chars {
		if _, ok := s2Chars[c]; ok {
			commomCharsCount += min(s2Chars[c], s1Chars[c])
		}
	}
	return commomCharsCount
}

func allChars(s string) map[byte]int {
	chars := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		chars[s[i]]++
	}
	return chars
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
