package main

func reverseInParentheses(input string) string {
	stack := make([]byte, 0, len(input))
	for i := 0; i < len(input); i++ {
		char := input[i]
		if char == ')' {
			tmp := make([]byte, 0, len(input))
			end := len(stack) - 1
			for stack[end] != '(' {
				tmp = append(tmp, stack[end])
				end--
			}
			stack = append(stack[:end], tmp...)
		} else {
			stack = append(stack, char)
		}
	}
	return string(stack)
}
