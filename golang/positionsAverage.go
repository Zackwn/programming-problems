package main

func PosAverage(str string) float64 {
	strs := split(str)
	n := len(strs[0])
	var com, all float64
	for i := 0; i < len(strs)-1; i++ {
		for j := i + 1; j < len(strs); j++ {
			for z := 0; z < n; z++ {
				if strs[j][z] == strs[i][z] {
					com++
				} else {
					all++
				}
			}
		}
	}
	return (com / (all + com)) * 100
}

func split(str string) []string {
	result := make([]string, 0)
	j, i := 0, 0
	for i <= len(str) {
		if i == len(str) || str[i] == ',' {
			result = append(result, str[j:i])
			i++
			j = i + 1
		}
		i++
	}
	return result
}
