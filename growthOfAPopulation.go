package main

func NbYear(p0 int, percent float64, aug int, p int) int {
	year := 0
	population := float64(p0)
	for int(population) < p {
		population += population * (percent / 100)
		population += float64(aug)
		year++
	}
	return year
}
