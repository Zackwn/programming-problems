package main

func almostIncreasingSequence(sequence []int) bool {
	for i := 0; i < len(sequence)-1; i++ {
		if sequence[i] >= sequence[i+1] {
			var ok bool
			ok = check(sequence, i)
			if ok {
				return true
			}
			ok = check(sequence, i+1)
			if ok {
				return true
			}
			break
		}
	}
	return false
}

func check(sequence []int, ig int) bool {
	a := make([]int, len(sequence)-1)
	ai := 0
	for i, e := range sequence {
		if i != ig {
			a[ai] = e
			ai++
		}
	}
	for i := 0; i < len(a)-1; i++ {
		if a[i] >= a[i+1] {
			return false
		}
	}
	return true
}
