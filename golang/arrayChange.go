package main

func arrayChange(inputArray []int) int {
	moves := 0
	for i := 1; i < len(inputArray); i++ {
		for inputArray[i-1] >= inputArray[i] {
			inputArray[i]++
			moves++
		}
	}
	return moves
}
