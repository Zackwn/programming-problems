package main

func RemovNb(m uint64) [][2]uint64 {
	var i uint64 = 1
	var sum uint64 = 0
	for i <= m {
		sum += i
		i++
	}
	i = 1
	result := [][2]uint64{}
	for i <= m {
		a := i
		b := (sum - a) / (a + 1)
		if b <= m && a*b == sum-a-b {
			result = append(result, [2]uint64{a, b})
		}
		i++
	}
	if len(result) == 0 {
		return nil
	}
	return result
}
