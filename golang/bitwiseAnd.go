package main

func bitwiseAnd(N int32, K int32) int32 {
	var result int32
	for a := int32(1); a <= N; a++ {
		for b := a + 1; b <= N; b++ {
			bitwise := a & b
			if bitwise < K && bitwise > result {
				result = bitwise
			}
		}
	}
	return result
}
