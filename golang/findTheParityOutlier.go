package main

func FindOutlier(ints []int) int {
	var even, odd *int
	for j, i := range ints {
		if even != nil && odd != nil {
			if i%2 == 0 {
				return *odd
			}
			return *even
		}
		if i%2 == 0 {
			even = &ints[j]
		} else {
			odd = &ints[j]
		}
	}
	return ints[len(ints)-1]
}
