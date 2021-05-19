package main

import "sort"

func InArray(array1 []string, array2 []string) []string {
	r := make([]string, 0, len(array1))
	i := 0
	for i < len(array1) {
		substr := array1[i]
	strloop:
		for _, str := range array2 {
			j := 0
			for j < len(str) {
				x := 0
				for x+j < len(str) && x < len(substr) && str[x+j] == substr[x] {
					x++
				}
				if x == len(substr) {
					for _, s := range r {
						if s == substr {
							break strloop
						}
					}
					r = append(r, substr)
					break strloop
				}
				j++
			}
		}
		i++
	}
	sort.Strings(r)
	return r
}
