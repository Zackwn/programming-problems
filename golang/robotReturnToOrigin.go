package main

func judgeCircle(moves string) bool {
	var x, y int
	for _, move := range moves {
		switch move {
		case 'U':
			y += 1
		case 'D':
			y -= 1
		case 'R':
			x += 1
		case 'L':
			x -= 1
		}
	}
	return x == 0 && y == 0
}
