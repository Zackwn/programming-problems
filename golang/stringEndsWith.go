package main

func solution(str, ending string) bool {
	if len(str) < len(ending) {
		return false
	}
	i := 0
	for i < len(ending) {
		stri := (len(str) - len(ending)) + i
		if str[stri] != ending[i] {
			return false
		}
		i++
	}
	return true
}
