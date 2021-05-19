package main

var MORSE_CODE map[string]string

func DecodeMorse(morseCode string) string {
	result := ""
	morse := []rune(morseCode)
	for morse[0] == ' ' {
		morse = morse[1:]
	}
	spacesInRow := 0
	morseChar := ""
	for i := 0; i < len(morse); i++ {
		if morse[i] == ' ' {
			spacesInRow++
			continue
		}
		if spacesInRow != 0 {
			result += MORSE_CODE[morseChar]
			morseChar = string(morse[i])
			if spacesInRow == 3 {
				result += " "
			}
		} else {
			morseChar += string(morse[i])
		}
		spacesInRow = 0
	}
	result += MORSE_CODE[morseChar]
	return result
}
