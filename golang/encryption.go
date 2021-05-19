package main

import (
	"math"
)

func encryption(s string) string {
	s = deleteSpaces(s)
	ssr := esqrt(float64(len(s)))
	rows := int(math.Floor(ssr))
	columns := int(math.Ceil(ssr))

	if rows == columns {
		rows--
	}

	trows := rows
	if rows*columns < len(s) {
		trows++
	}

	result := ""
	for i := 0; i < columns; i++ {
		r := 0
		for k := 0; k < trows; k++ {
			if k+r+i >= len(s) {
				break
			}
			result += string(s[k+r+i])
			r += rows
		}
		result += " "
	}
	return result
}

func esqrt(x float64) float64 {
	var lz float64
	z := 1.0
	for i := 0; i < 1000; i++ {
		lz = z
		z -= (z*z - x) / (2 * z)
		if lz == z {
			return z
		}
	}
	return z
}

func deleteSpaces(s string) string {
	result := ""
	for index := 0; index < len(s); index++ {
		if s[index] != ' ' {
			result += string(s[index])
		}
	}
	return result
}
