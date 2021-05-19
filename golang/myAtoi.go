package main

import (
	"math"
	"unicode"
)

func myAtoi(s string) int {
	isUnsigned := false
	signAlreadyDefined := false
	var integer int
	for index, char := range s {
		if char == ' ' {
			if signAlreadyDefined {
				break
			}
			continue
		}
		if char == '-' || char == '+' {
			if signAlreadyDefined {
				break
			}
			isUnsigned = char == '-'
			signAlreadyDefined = true
			continue
		}
		if unicode.IsDigit(char) {
			for i := index; i < len(s); i++ {
				if unicode.IsDigit(rune(s[i])) {
					// '0' = 48
					integer = integer*10 + int(rune(s[i])-'0')
					if integer > math.MaxInt32 {
						if isUnsigned {
							return math.MinInt32
						} else {
							return math.MaxInt32
						}
					}
				} else {
					break
				}
			}
			break
		} else {
			break
		}
	}
	if isUnsigned {
		integer *= -1
	}
	return integer
}
