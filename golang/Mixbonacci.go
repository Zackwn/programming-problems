package main

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func pad(n int) int {
	if n == 0 {
		return 1
	} else if n < 3 {
		return 0
	}
	return pad(n-2) + pad(n-3)
}

func jac(n int) int {
	if n < 2 {
		return n
	}
	return jac(n-1) + 2*jac(n-2)
}

func pel(n int) int {
	if n < 2 {
		return n
	}
	return 2*pel(n-1) + pel(n-2)
}

func tri(n int) int {
	if n < 2 {
		return 0
	} else if n == 2 {
		return 1
	}
	return tri(n-1) + tri(n-2) + tri(n-3)
}

func tet(n int) int {
	if n < 3 {
		return 0
	} else if n == 3 {
		return 1
	}
	return tet(n-1) + tet(n-2) + tet(n-3) + tet(n-4)
}

func Mixbonacci(pattern []string, length int) []int64 {
	if len(pattern) == 0 || length == 0 {
		return []int64{}
	}
	var result = make([]int64, length)

	nacciCounts := map[string]int{}
	nacciCounts["fib"] = 0
	nacciCounts["pad"] = 0
	nacciCounts["jac"] = 0
	nacciCounts["pel"] = 0
	nacciCounts["tri"] = 0
	nacciCounts["tet"] = 0

	nacciFuncs := map[string]func(n int) int{}
	nacciFuncs["fib"] = fib
	nacciFuncs["pad"] = pad
	nacciFuncs["jac"] = jac
	nacciFuncs["pel"] = pel
	nacciFuncs["tri"] = tri
	nacciFuncs["tet"] = tet

	i, patternIndex := 0, 0
	for i < length {
		if patternIndex == len(pattern) {
			patternIndex = 0
		}
		p := pattern[patternIndex]
		result[i] = int64(nacciFuncs[p](nacciCounts[p]))
		nacciCounts[p]++
		i++
		patternIndex++
	}
	return result
}
