package main

import "strconv"

func Solution(list []int) string {
	result := ""
	i := 0
	for i < len(list) {
		l, r := i, i+1
		for r < len(list) {
			if (list[l] + 1) == list[r] {
				l++
				r++
				continue
			}
			if (list[l] - 1) == list[r] {
				l++
				r++
				continue
			}
			break
		}
		if list[i]-list[r-1] == -1 {
			result += "," + strconv.Itoa(list[i]) + "," + strconv.Itoa(list[r-1])
		} else if list[i] == list[r-1] {
			result += "," + strconv.Itoa(list[i])
		} else {
			result += "," + strconv.Itoa(list[i]) + "-" + strconv.Itoa(list[r-1])
		}
		i = r
	}
	return result[1:]
}
