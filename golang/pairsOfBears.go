package main

func CheckPairs(nbBear int, bears string) (couples string, hasEnoughCouple bool) {
	var male, female byte = 'B', '8'
	x := 0

	for i := 0; i < len(bears)-1; i++ {
		if (bears[i] == male && bears[i+1] == female) ||
			(bears[i] == female && bears[i+1] == male) {
			couples += bears[i : i+2]
			i++
			x++
		}
	}

	hasEnoughCouple = x*2 >= nbBear

	return
}
