package main

func rotateStr(s string) string {
	return s[1:] + string(s[0])
}

func reverseStr(str string) string {
	s := []rune(str)
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	return string(s)
}

func Revrot(s string, n int) string {
	// your code
	if n <= 0 || len(s) == 0 || len(s) < n {
		return ""
	}
	result := ""
	i := 0
	for i < len(s) {
		chunksum := 0
		chunk := ""
		j := 0
		for i < len(s) && j < n {
			chunk += string(s[i])
			digit := int(s[i] - 48)
			chunksum += (digit * digit) * digit
			i++
			j++
		}
		if j != n {
			break
		}
		if chunksum%2 == 0 {
			result += reverseStr(chunk)
		} else {
			result += rotateStr(chunk)
		}
	}
	return result
}
