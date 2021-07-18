package main

func sqrt(x float64) float64 {
	var lz float64
	z := 1.0
	for i := 0; i < 1000; i++ {
		lz = z
		z -= (z*z - x) / (2 * z)
		if lz == z {
			return z
		}
	}
	return z
}

func solve(n int) (int, bool) {
	sum := 0
	sqrtRoot := int(sqrt(float64(n)))
	for i := 1; i <= sqrtRoot; i++ {
		if n%i == 0 {
			sum += i * i
			if n/i != i {
				sum += (n / i) * (n / i)
			}
		}
	}
	sumSR := sqrt(float64(sum))
	return sum, sumSR == float64(int(sumSR))
}

func ListSquared(m, n int) [][]int {
	result := make([][]int, 0)
	for m <= n {
		sum, valid := solve(m)
		if valid {
			result = append(result, []int{m, sum})
		}
		m++
	}
	return result
}
