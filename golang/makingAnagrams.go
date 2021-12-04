package main

func makingAnagrams(s1 string, s2 string) int32 {
	var deletions int32
	s1c, s2c := make(map[rune]int32), make(map[rune]int32)
	for _, c := range s1 {
		s1c[c]++
	}
	for _, c := range s2 {
		s2c[c]++
	}
	for char, count := range s1c {
		count2 := s2c[char]
		if count2 != count {
			deletions += diff(count2, count)
			delete(s2c, char) // to avoid counting char again
		}
	}
	for char, count := range s2c {
		count2 := s1c[char]
		if count2 != count {
			deletions += diff(count2, count)
		}
	}
	return deletions
}

func diff(a, b int32) int32 {
	if a > b {
		return a - b
	}
	return b - a
}
