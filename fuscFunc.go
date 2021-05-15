package kata

func Fusc(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	if n%2 == 0 {
		return Fusc(n / 2)
	}
	return Fusc(n/2) + Fusc(n/2+1)
}
