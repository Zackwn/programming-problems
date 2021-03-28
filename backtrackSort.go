func Sort(slice []int32) {
	ahead, back := 1, 0

	swap := makeSwapper(slice)

	for ahead < len(slice) {
		if slice[ahead] < slice[back] {
			swap(ahead, back)

			// backtrack
			for back-1 >= 0 && slice[back] < slice[back-1] {
				swap(back, back-1)
				back--
			}
			back = ahead - 1
		}
		ahead++
		back++
	}
}

func makeSwapper(slice []int32) func(int, int) {
	return func(i1, i2 int) {
		slice[i1], slice[i2] = slice[i2], slice[i1]
	}
}