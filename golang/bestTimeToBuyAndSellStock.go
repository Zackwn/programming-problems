package main

import "fmt"

func main() {
	prices := []int{1, 2, 3, 4, 5}
	fmt.Println(BestTimeToBuyAndSellStock(prices)) // 4
	prices = []int{}
	fmt.Println(BestTimeToBuyAndSellStock(prices)) // 0
	prices = []int{2, 2, 2}
	fmt.Println(BestTimeToBuyAndSellStock(prices)) // 0
	prices = []int{3, 5, 2, 9, 10, 7}
	fmt.Println(BestTimeToBuyAndSellStock(prices)) // 8
	prices = []int{4, 5, 2, 1, 6, 10, 4, 9, 11}
	fmt.Println(BestTimeToBuyAndSellStock(prices)) // 10

}

func BestTimeToBuyAndSellStock(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min := prices[0]
	max := prices[0]
	for _, n := range prices {
		if min > n {
			min = n
		}
		if max < n {
			max = n
		}
	}
	profit := max - min
	return profit
}
