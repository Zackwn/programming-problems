package main

import "fmt"

func TestArrayWithOddSum() {
	a := []int{2, 3}
	fmt.Println(ArrayWithOddSum(a)) // YES
	a = []int{2, 2, 8, 8}
	fmt.Println(ArrayWithOddSum(a)) // NO
	a = []int{3, 3, 3}
	fmt.Println(ArrayWithOddSum(a)) // YES
	a = []int{5, 5, 5, 5}
	fmt.Println(ArrayWithOddSum(a)) // NO
	a = []int{1, 1, 1, 1}
	fmt.Println(ArrayWithOddSum(a)) // NO
}

// return "YES" or "NO"
func ArrayWithOddSum(a []int) string {
	oddNums := 0
	for i := 0; i < len(a); i++ {
		if a[i]%2 != 0 {
			oddNums++
		}
	}

	if oddNums%2 == 0 {
		return "NO"
	} else {
		return "YES"
	}
}
