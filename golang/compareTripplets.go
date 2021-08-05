package main

func compareTriplets(a []int32, b []int32) []int32 {
	scores := make([]int32, 2)
	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			scores[0]++
		} else if b[i] > a[i] {
			scores[1]++
		}
	}
	return scores

}
