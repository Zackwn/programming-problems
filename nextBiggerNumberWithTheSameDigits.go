package kata

func NextBigger(n int) int {
	if n/10 == 0 {
		return -1
	}
	digit := n % 10
	n = n / 10
	prevdigit := n % 10
	passedigits := -1
	pm := 1
	for n > 0 {
		if prevdigit >= digit {
			if prevdigit != digit && digit > (n/10)%10 {
				pm = pm * 10
				passedigits = passedigits*10 + n%10
				n = n / 10
				prevdigit = n % 10
				continue
			}
			if passedigits == -1 {
				passedigits = digit
			} else {
				passedigits = passedigits*10 + digit
			}
			pm = pm * 10
			digit = prevdigit
			n = n / 10
			prevdigit = n % 10
		} else {
			n = n / 10
			n = n*10 + digit
			n = n*10 + prevdigit
			if passedigits != -1 {
				pp := 1
				pdt := passedigits
				for pdt >= 0 {
					if n%10 > pdt%10 {
						temp := n % 10
						n = (n/10)*10 + pdt%10
						passedigits = (temp * pp) + (passedigits % pp)
						break
					}
					pp = pp * 10
					pdt = pdt / 10
				}
				n = (n * pm) + passedigits
			}
			return n
		}
	}
	return -1
}
