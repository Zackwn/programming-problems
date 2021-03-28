func cutTheSticks(arr []int32) []int32 {
	sort(arr)
	lengths := []int32{}

	for len(arr) > 0 {
		lengths = append(lengths, int32(len(arr)))

		value := arr[0]
		i := 0
		for i < len(arr) {
			if arr[i] == value {
				arr = append(arr[:i], arr[i+1:]...)
			} else {
				arr[i] -= value
				i++
			}
		}
	}
	return lengths
}

func sort(arr []int32) {
	for index, element := range arr {
		smallerElementIndex := -1
		for searchIndex := index + 1; searchIndex < len(arr); searchIndex++ {
			if element > arr[searchIndex] {
				if smallerElementIndex == -1 {
					smallerElementIndex = searchIndex
				} else {
					if arr[smallerElementIndex] > arr[searchIndex] {
						smallerElementIndex = searchIndex
					}
				}
			}
		}
		if smallerElementIndex != -1 {
			arr[index] = arr[smallerElementIndex]
			arr[smallerElementIndex] = element
		}
	}
}