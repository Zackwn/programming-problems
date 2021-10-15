package main

func Parse(data string) []int {
	value, result := 0, make([]int, 0)
	for _, c := range data {
		switch c {
		case 'i':
			value++
		case 'd':
			value--
		case 's':
			value *= value
		case 'o':
			result = append(result, value)
		}
	}
	return result
}
