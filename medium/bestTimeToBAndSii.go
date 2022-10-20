package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 7, 4}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	profit := 0
	buyPrice := prices[0]
	count := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			count++
		} else {
			if count > 0 {
				profit += prices[i-1] - buyPrice
			}
			buyPrice = prices[i]
			count = 0
		}
	}

	if count > 0 {
		profit += prices[len(prices)-1] - buyPrice
	}

	return profit
}

// func maxProfit(prices []int) int {
// 	if len(prices) == 0 {
// 		return 0
// 	}
// 	profit := 0
// 	for i := 1; i < len(prices); i++ {
// 		if prices[i] > prices[i-1] {
// 			profit += prices[i] - prices[i-1]
// 		}
// 	}
// 	return profit
// }
