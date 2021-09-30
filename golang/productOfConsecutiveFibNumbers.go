package main

func ProductFib(prod uint64) [3]uint64 {
	var left, right, found uint64 = 0, 1, 0
	p := left * right
	for p <= prod {
		if prod == p {
			found = 1
			break
		} else {
			left, right = right, left+right
			p = left * right
		}
	}
	return [3]uint64{left, right, found}
}
