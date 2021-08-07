package main

func matrixElementsSum(matrix [][]int) int {
	var sum int
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			if checkRoom(matrix, x, y) {
				sum += matrix[y][x]
			}
		}
	}
	return sum
}

func checkRoom(matrix [][]int, x, y int) bool {
	if y > 0 {
		ok := matrix[y-1][x] != 0
		if !ok {
			return false
		}
		return checkRoom(matrix, x, y-1)
	}
	return true
}
