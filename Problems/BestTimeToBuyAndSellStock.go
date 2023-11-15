package main

// Given an array prices where prices[i] is the price of a given stock on the ith day
// Find the maximum profit you can achieve, return 0 if you can't make a profit

// Solution: Use two variables, minPrice and maxProfit
// minPrice is the minimum price of the stock
// maxProfit is the maximum profit you can achieve
// If the current price is less than minPrice, then you can buy the stock at the current price
// If the difference between the current price and minPrice is greater than maxProfit, then you can sell the stock at the current price
func maxProfit(prices []int) int {
	var minPrice, maxProfit int
	for i := 0; i < len(prices); i++ {
		if i == 0 {
			minPrice = prices[i]
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}
