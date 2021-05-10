package kata

import (
	"strconv"
)

func Decomp(n int) string {
	primes := []int{2}
PrimeLoop:
	for i := 3; i <= n; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				continue PrimeLoop
			}
		}
		primes = append(primes, i)
	}

	totalprimes := []int{}
	for n > 1 {
		num := n
		i := 0
		for i < len(primes) {
			if num%primes[i] == 0 {
				totalprimes = append(totalprimes, primes[i])
				num = num / primes[i]
			} else {
				i++
			}
		}
		n--
	}

	decomp := ""
	for _, prime := range primes {
		count := 0
		for _, p := range totalprimes {
			if prime == p {
				count++
			}
		}
		if count > 0 {
			if len(decomp) != 0 {
				decomp += " * "
			}
			if count == 1 {
				decomp += strconv.Itoa(prime)
			} else {
				decomp += strconv.Itoa(prime) + "^" + strconv.Itoa(count)
			}
		}
	}
	return decomp
}
