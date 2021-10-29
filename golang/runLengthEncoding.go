package main

func RunLengthEncode(input string) string {
	encoded := ""
	for i := 0; i < len(input); i++ {
		b, count := input[i], 1
		j := i + 1
		for j < len(input) && input[j] == b {
			count++
			j++
		}
		i = j - 1
		if count > 1 {
			encoded += String(count)
		}
		encoded += string(b)
	}
	return encoded
}

func RunLengthDecode(input string) string {
	decoded := ""
	for i := 0; i < len(input); i++ {
		count := 0
		j := i
		for j < len(input) && digit(input[j]) {
			count = count*10 + Int(input[j])
			j++
		}
		i = j
		if count == 0 {
			count = 1
		}
		for count > 0 {
			decoded += string(input[j])
			count--
		}
	}
	return decoded
}

func digit(b byte) bool {
	return b >= '0' && b <= '9'
}

func String(n int) string {
	r := ""
	for n > 0 {
		r = string(rune(n%10+'0')) + r
		n /= 10
	}
	return r
}

func Int(b byte) int {
	return int(b - '0')
}
