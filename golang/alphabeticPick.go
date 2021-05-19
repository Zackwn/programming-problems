package main

func alphabeticPick(s1, s2 string) string {
	if s1 == s2 {
		return s1
	}

	i := 0
	for i < len(s1) && i < len(s2) {
		if s1[i] > s2[i] {
			return s2
		} else if s1[i] < s2[i] {
			return s1
		} else {
			// s1[i] == s2[i]
			i++
		}
	}

	if len(s1) > len(s2) {
		return s2
	} else {
		return s1
	}
}
