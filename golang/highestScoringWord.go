package main

func High(s string) string {
	hw, hc := "", 0
	for i := 0; i < len(s); i++ {
		j := i + 1
		for j < len(s) && s[j] != ' ' {
			j++
		}
		w, c := s[i:j], 0
		z := i
		for z < j {
			c += int(s[z]) - 'a' + 1
			z++
		}
		if c > hc {
			hc = c
			hw = w
		}
		i = j
	}
	return hw
}
