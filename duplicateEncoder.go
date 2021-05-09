package main

import "strings"

func DuplicateEncode(word string) string {
	word = strings.ToLower(word)
	words := map[rune]string{}
	for _, char := range word {
		_, ok := words[char]
		if ok {
			words[char] = ")"
		} else {
			words[char] = "("
		}
	}
	result := ""
	for _, c := range word {
		v := words[c]
		result += v
	}
	return result
}
