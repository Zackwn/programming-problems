func isValid(s string) bool {
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	openBrackets := []rune{}
	for _, char := range s {
		expectedOpenBracket, has := bracketMap[char]
		if has {
			if len(openBrackets) > 0 {
				index := len(openBrackets) - 1
				if openBrackets[index] == expectedOpenBracket {
					openBrackets = append(openBrackets[:index], openBrackets[index+1:]...)
					continue
				}
			}
			// the slice is empty or the opening is wrong
			return false
		} else {
			openBrackets = append(openBrackets, char)
		}
	}
	return len(openBrackets) <= 0
}