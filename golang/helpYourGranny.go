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

func Tour(arrFriends []string, ftwns map[string]string, h map[string]float64) int {
	r := h[ftwns[arrFriends[0]]]
	i := 1
	for i < len(arrFriends) {
		ltd := h[ftwns[arrFriends[i-1]]]  // last town distance
		td, ok := h[ftwns[arrFriends[i]]] // town distance
		if ok == false {
			break
		}
		d := sqrt((td * td) - (ltd * ltd))
		r += d
		i++
	}
	r += h[ftwns[arrFriends[i-1]]]
	return int(r)
}
