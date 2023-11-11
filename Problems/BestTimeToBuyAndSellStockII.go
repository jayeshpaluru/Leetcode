package main

// Given an int array prices where prices[i] is the price of a given stock on the ith day
// You can only hold onto one stock at a time
// Find the maximum profit you can achieve, return 0 if you can't make a profit
func maxProfitII(prices []int) int {
	var maxProfit int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit = maxProfit + prices[i] - prices[i-1]
		}
	}
	return maxProfit
}
