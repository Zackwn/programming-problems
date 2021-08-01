package main

func makeArrayConsecutive2(statues []int) int {
	sttsMap := make(map[int]bool)
	min, max := 20, 0
	for _, stt := range statues {
		sttsMap[stt] = true
		if stt > max {
			max = stt
		}
		if stt < min {
			min = stt
		}
	}
	r := 0
	for min < max {
		_, ok := sttsMap[min]
		if !ok {
			r++
		}
		min++
	}
	return r
}
