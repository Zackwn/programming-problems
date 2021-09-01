package main

func arrayMaximalAdjacentDifference(inputArray []int) int {
	max := 0
	for i := 0; i < len(inputArray)-1; i++ {
		d := diff(inputArray[i], inputArray[i+1])
		if d > max {
			max = d
		}
	}
	return max
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
