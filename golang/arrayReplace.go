package main

func arrayReplace(a []int, r int, s int) []int {
	for i, e := range a {
		if e == r {
			a[i] = s
		}
	}
	return a
}
