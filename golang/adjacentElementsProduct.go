package main

func adjacentElementsProduct(inputArray []int) int {
	max := -2147483648
	for i := 0; i < len(inputArray)-1; i++ {
		p := inputArray[i] * inputArray[i+1]
		if p > max {
			max = p
		}
	}
	return max
}
