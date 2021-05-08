package main

func toWeirdCase(str string) string {
	// Your code here and happy coding!
	s := []rune(str)
	i := 0
	for i < len(s) {
		j := 0
		for i+j < len(s) && s[i+j] != ' ' {
			if j%2 == 0 && s[i+j] >= 97 && s[i+j] <= 122 {
				// to upper case
				s[i+j] = s[i+j] - 32
			} else if j%2 != 0 && s[i+j] >= 65 && s[i+j] <= 90 {
				// to lower case
				s[i+j] = s[i+j] + 32
			}
			j++
		}
		i += j + 1 // +1 for skiping space
	}
	return string(s)
}
