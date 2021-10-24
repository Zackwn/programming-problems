package main

import (
	"sort"
)

type ByModulo []string

func (arr ByModulo) Len() int {
	return len(arr)
}
func (arr ByModulo) Less(i, j int) bool {
	// return bymodulo[i] < bymodulo[j]
	a, b := len(arr[i])%3, len(arr[j])%3
	if a == b {
		index := 0
		for index < len(arr[i]) && index < len(arr[j]) && arr[i][index] == arr[j][index] {
			index++
		}
		if index < len(arr[i]) && index < len(arr[j]) {
			return arr[i][index] < arr[j][index]
		} else {
			return len(arr[i]) < len(arr[j])
		}
	}
	return a < b
}

func (arr ByModulo) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func RemainderSorting(arr []string) []string {
	sort.Sort(ByModulo(arr))
	return arr
}
