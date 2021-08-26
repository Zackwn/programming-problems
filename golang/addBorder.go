package main

func addBorder(picture []string) []string {
	l := len(picture[0]) + 2
	result := make([]string, len(picture)+2)
	result[0] = ast(l)
	i := 1
	for i < len(result)-1 {
		result[i] = "*" + picture[i-1] + "*"
		i++
	}
	result[len(result)-1] = ast(l)
	return result
}

func ast(n int) string {
	r := ""
	for n > 0 {
		r += "*"
		n--
	}
	return r
}
