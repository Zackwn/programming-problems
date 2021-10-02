package main

func Anagrams(word string, words []string) []string {
	wcc, lwcc, result := make(map[rune]int), make(map[rune]int), make([]string, 0)
	for _, c := range word {
		wcc[c]++
	}
MLoop:
	for _, w := range words {
		// clean map
		lwcc = make(map[rune]int)
		for _, c := range w {
			lwcc[c]++
		}
		if len(lwcc) != len(wcc) {
			continue MLoop
		}
		for key, count := range wcc {
			if lwcc[key] != count {
				continue MLoop
			}
		}
		result = append(result, w)
	}
	if len(result) == 0 {
		return nil
	}
	return result
}
