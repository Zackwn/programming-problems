package kata

func JosephusSurvivor(n, k int) int {
	//Your code:
	people := []int{}
	for j := 1; j <= n; j++ {
		people = append(people, j)
	}
	index := 0
	for len(people) != 1 {
		index += k
		for index > len(people) {
			index -= len(people)
		}
		i := index - 1
		people = append(people[:i], people[i+1:]...)
		index--
	}
	return people[0]
}
