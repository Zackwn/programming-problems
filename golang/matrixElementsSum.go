package main

func matrixElementsSum(matrix [][]int) int {
	var sum int
	for x := 0; x < len(matrix[0]); x++ {
		for y := 0; y < len(matrix); y++ {
			if matrix[y][x] == 0 {
				break
			}
			sum += matrix[y][x]
		}
	}
	return sum
}
