package main

func Solution(ar []int) int {
	arr := make([]int, len(ar))
	copy(arr, ar)
	min := min(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] != min {
			if arr[i]%min != 0 {
				arr[i] = arr[i] % min
				if min > arr[i] {
					min = arr[i]
				}
			} else {
				arr[i] = 0
			}
		}
		if arr[i] == 1 {
			return len(arr)
		}
	}
	return min * len(arr)
}

func min(arr []int) int {
	m := arr[0]
	for i := 1; i < len(arr); i++ {
		if m > arr[i] {
			m = arr[i]
		}
	}
	return m
}
