package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 6, 1, 2, 9}
	sort(arr)
	fmt.Printf("%v\n", arr)
}

func sort(arr []int) []int {
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
			fmt.Printf("change: element(%v) by %v\n", element, arr[index])
		}
	}

	return arr
}
