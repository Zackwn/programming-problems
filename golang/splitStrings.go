package main

func Solution(str string) []string {
	pairs := make([]string, 0, len(str)/2)
	for i := 0; i < len(str); i += 2 {
		s, e := i, i+2
		if e > len(str) {
			e = len(str)
		}
		pair := str[s:e]
		if len(pair) != 2 {
			pair += "_"
		}
		pairs = append(pairs, pair)
	}
	return pairs
}
