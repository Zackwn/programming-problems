package main

import "strconv"

func EncryptThis(text string) string {
	if text == "" {
		return text
	}
	encrypt := ""
	i := 0
	for i < len(text) {
		c := text[i]
		// add ascii code
		encrypt += strconv.Itoa(int(c))
		s, e := i, i+1
		for e < len(text) && text[e] != ' ' {
			e++
		}
		if e-2 > s {
			// switch and add letters
			encrypt += string(text[e-1])
			encrypt += text[s+2 : e-1]
			encrypt += string(text[s+1])
		} else if e-2 == s {
			// add letter (ex: in - add 'n')
			encrypt += string(text[e-1])
		}
		encrypt += " "
		i = e + 1
	}
	// remove space in the end
	return encrypt[:len(encrypt)-1]
}
