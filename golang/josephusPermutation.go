package main

func Josephus(items []interface{}, k int) []interface{} {
	result := []interface{}{}
	index := 0
	for len(items) != 0 {
		index += k
		for index > len(items) {
			index -= len(items)
		}
		i := index - 1
		result = append(result, items[i])
		items = append(items[:i], items[i+1:]...)
		index--
	}
	return result
}
