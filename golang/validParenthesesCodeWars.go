package main

func ValidParentheses(parens string) bool {
	opening := 0
	for _, char := range parens {
		if char == ')' {
			if opening > 0 {
				opening--
				continue
			}
			// misplaced ')'
			return false
		} else {
			opening++
		}
	}
	return opening == 0
}
