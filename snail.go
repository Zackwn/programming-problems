package kata

func Snail(snaipMap [][]int) []int {
	if len(snaipMap) == 0 {
		return []int{}
	}
	result := make([]int, 0, len(snaipMap)*len(snaipMap[0]))
	k := 0
	vertical := 0
	horizontal := 0
	for k < len(snaipMap)/2+1 {
		// left to right
		for horizontal < len(snaipMap[vertical])-k {
			result = append(result, snaipMap[vertical][horizontal])
			horizontal++
		}
		horizontal--
		vertical++
		// top to bottom
		for vertical < len(snaipMap)-k {
			result = append(result, snaipMap[vertical][horizontal])
			vertical++
		}
		vertical--
		horizontal--
		// right to left
		for horizontal >= 0+k {
			result = append(result, snaipMap[vertical][horizontal])
			horizontal--
		}
		horizontal++
		vertical--
		// bottom to top
		for vertical > k {
			result = append(result, snaipMap[vertical][horizontal])
			vertical--
		}
		vertical++
		horizontal++
		k++
	}
	return result
}
