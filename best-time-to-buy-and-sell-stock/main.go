package main

import (
	"fmt"
)

// leetcode ref: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
func main() {
	fmt.Printf("%d\n", maxProfit([]int{7, 1, 5, 3, 6, 4})) // output 5
	fmt.Printf("%d\n", maxProfit([]int{7, 6, 4, 3, 1}))    // output 0
}

func maxProfit(prices []int) int {
	// Track the minimum price seen so far (initially the first price)
	minPrice := prices[0]

	// Initialize maximum profit to 0
	maxProfit := 0

	// Iterate through each price in the array
	for i := 0; i < len(prices); i++ {
		// Update maximum profit if selling at current price yields higher profit
		// Profit = current price - minimum price seen so far
		remainder := prices[i] - minPrice
		if remainder > maxProfit {
			maxProfit = remainder
		}

		// Update minimum price if current price is lower
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}

	// Return the maximum profit achievable
	return maxProfit
}

func maxProfit2(prices []int) int {
	// Track the minimum price seen so far (initially the first price)
	minPrice := prices[0]

	// Initialize maximum profit to 0
	maxProfit := 0

	for _, price := range prices {
		// Update maximum profit if selling at current price yields higher profit
		// Profit = current price - minimum price seen so far

		maxProfit = max(maxProfit, price-minPrice)

		// Update minimum price if current price is lower

		minPrice = min(minPrice, price)
	}

	// Return the maximum profit achievable
	return maxProfit
}
