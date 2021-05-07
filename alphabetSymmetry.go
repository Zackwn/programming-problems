package main

func solve(slice []string) []int {
	// Your code here and happy coding!
	r := make([]int, 0, len(slice))
	for _, str := range slice {
		t := 0
		for position, char := range str {
			if (int(char-97) == position) || (int(char-65) == position) {
				t++
			}
		}
		r = append(r, t)
	}
	return r
}
