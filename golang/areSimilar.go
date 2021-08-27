package main

func areSimilar(a []int, b []int) bool {
	swaps := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			if swaps >= 1 {
				return false
			}
			ok := false
			for n := 0; n < len(b); n++ {
				if b[n] == a[i] {
					if b[i] != a[n] {
						continue
					}
					b[i], b[n] = b[n], b[i]
					swaps++
					ok = true
					break
				}
			}
			if !ok {
				return false
			}
		}
	}
	return true
}
