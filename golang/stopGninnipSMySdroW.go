package main

func SpinWords(str string) (spinned string) {
	spin := func(s string) string {
		if len(s) < 5 {
			return s
		}
		b := []byte(s)
		for j, i := 0, len(b)-1; j < i; i, j = i-1, j+1 {
			b[j], b[i] = b[i], b[j]
		}
		return string(b)
	}
	j, i := 0, 0
	for i < len(str) {
		if str[i] == ' ' {
			spinned += spin(str[j:i]) + " "
			j = i + 1
		}
		i++
	}
	spinned += spin(str[j:])
	return
}
