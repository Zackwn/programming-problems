package main

func FindShort(s string) int {
	// your code
	shortest := 99999
	i := 0
	for i < len(s) {
		length := 0
		for i < len(s) && s[i] != ' ' {
			length++
			i++
		}
		if shortest > length {
			shortest = length
		}
		i++
	}
	return shortest
}
