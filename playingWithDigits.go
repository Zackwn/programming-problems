package main

import "strconv"

func Pow(base, exponent int) int {
	mult := base
	for i := 1; i < exponent; i++ {
		base = base * mult
	}
	return base
}

func DigPow(n, p int) int {
	nlen := len(strconv.Itoa(n))
	bases := make([]int, 0, nlen)

	tn := n
	for tn > 0 {
		bases = append(bases, tn%10)
		tn = tn / 10
	}

	x := 0
	for i := nlen - 1; i >= 0; i-- {
		x += Pow(bases[i], p)
		p++
	}

	f := 1
	for n*f <= x {
		if n*f == x {
			return f
		}
		f++
	}
	return -1
}
