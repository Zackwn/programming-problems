package kata

import (
	"math"
)

func toFixed(num float64) float64 {
	output := math.Pow(10, 6.0)
	return float64(int(num*output)) / output
}

func Going(n int) float64 {
	result, acc := 1.0, 1.0
	for n > 1 {
		acc = acc / float64(n)
		result += acc
		n--
	}
	return toFixed(result)
}
