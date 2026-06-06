package main

import (
	"fmt"
)

// leetcode ref: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/
func main() {
	fmt.Printf("%d\n", maxProfit([]int{7, 1, 5, 3, 6, 4})) // output 7
	fmt.Printf("%d\n", maxProfit([]int{1, 2, 3, 4, 5}))    // output 4
	fmt.Printf("%d\n", maxProfit([]int{7, 6, 4, 3, 1}))    // output 0
}

func maxProfit(prices []int) int {
	maxProfit := 0

	for i, price := range prices {
		if i+1 == len(prices) {
			break
		}

		if (prices[i+1] - price) > 0 {
			maxProfit += prices[i+1] - price
		}
	}

	return maxProfit
}
