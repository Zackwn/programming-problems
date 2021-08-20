package main

func alternatingSums(a []int) []int {
	team1, team2 := 0, 0
	for i := 0; i < len(a); i++ {
		if i%2 == 0 {
			team1 += a[i]
		} else {
			team2 += a[i]
		}
	}
	return []int{team1, team2}
}
