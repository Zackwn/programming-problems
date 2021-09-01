package main

func palindromeRearranging(inputString string) bool {
	var chars = make(map[rune]int)
	for _, c := range inputString {
		chars[c]++
	}
	oddCount := 0
	for _, charCount := range chars {
		if charCount%2 != 0 {
			oddCount++
		}
	}
	return oddCount <= 1
}
