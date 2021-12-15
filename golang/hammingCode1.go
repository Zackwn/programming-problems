package main

func Encode(text string) string {
	var bits string

	for _, c := range text {
		// convert ASCII to binary
		d, b := int(c), ""
		for d > 0 {
			b = string(rune(d%2+'0')) + b
			d /= 2
		}

		// fix length
		for len(b) < 8 {
			b = "0" + b
		}

		// triple bits
		bit := ""
		for _, c := range b {
			for i := 0; i < 3; i++ {
				bit += string(c)
			}
		}

		bits += bit
	}

	return bits
}

func Decode(bits string) string {
	var text string

	n := 8 * 3
	for i := 0; i < len(bits); i += n {
		tbit := bits[i : i+n]

		// correct bits
		b := ""
		for j := 0; j < len(tbit); j += 3 {
			corrBit := "0"
			count := make(map[rune]int, 2)
			for _, c := range tbit[j : j+3] {
				count[c]++
			}
			if count['1'] > count['0'] {
				corrBit = "1"
			}
			b += corrBit
		}

		// convert binary to ASCII
		ascii := 0
		x := 1
		for i := len(b) - 1; i >= 0; i-- {
			ascii += int(b[i]-'0') * x
			x *= 2
		}

		text += string(rune(ascii))
	}

	return text
}
