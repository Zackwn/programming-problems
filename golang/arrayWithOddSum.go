package main

// import "fmt"

// func main() {
// 	var a []int
// 	n := 0
// 	x := 0
// 	fmt.Scan(&n)
// 	for i := 0; i < n; i++ {
// 		fmt.Scan(&x)
// 		a = make([]int, x)
// 		k := 0
// 		for j := 0; j < x; j++ {
// 			fmt.Scan(&k)
// 			a[j] = k
// 		}
// 		fmt.Println(ArrayWithOddSum(a))
// 	}
// }

func ArrayWithOddSum(a []int) string {
	oddNums := 0
	evenNums := 0
	for i := 0; i < len(a); i++ {
		if a[i]%2 == 0 {
			evenNums++
		} else {
			oddNums++
		}
	}

	if oddNums != 0 {
		if oddNums%2 == 0 && evenNums == 0 {
			return "NO"
		}
		return "YES"
	} else {
		return "NO"
	}
}
