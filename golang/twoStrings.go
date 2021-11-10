package main

func twoStrings(s1 string, s2 string) string {
	s1cMap := make(map[rune]bool)
	s2cMap := make(map[rune]bool)
	for _, c := range s1 {
		s1cMap[c] = true
	}
	for _, c := range s2 {
		s2cMap[c] = true
	}
	for s1Char := range s1cMap {
		_, ok := s2cMap[s1Char]
		if ok {
			return "YES"
		}
	}
	return "NO"
}
