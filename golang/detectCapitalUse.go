package main

func detectCapitalUse(word string) bool {
	var (
		hasUpper = false
		hasLower = false
	)
	for i, char := range word {
		if hasLower && hasUpper {
			return false
		}
		if char >= 'a' && char <= 'z' {
			hasLower = true
		} else if i != 0 && char >= 'A' && char <= 'Z' {
			hasUpper = true
		}
	}
	return !(hasLower && hasUpper)
}
