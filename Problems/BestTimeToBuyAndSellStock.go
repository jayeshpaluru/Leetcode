package main

// Given an array prices where prices[i] is the price of a given stock on the ith day
// Find the maximum profit you can achieve, return 0 if you can't make a profit

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
