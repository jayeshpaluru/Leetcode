package main

// There are n children standing in a line. Each child is assigned a rating value given in the integer array ratings.
// You are giving candies to these children subjected to the following requirements:
// Each child must have at least one candy.
// Children with a higher rating get more candies than their neighbors.
// Return the minimum number of candies you need to have to distribute the candies to the children.

func candy(ratings []int) int {
	n := len(ratings)
	if n == 1 {
		return 1
	}
	candies := make([]int, n)
	candies[0] = 1
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		} else {
			candies[i] = 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}
	sum := 0
	for _, v := range candies {
		sum += v
	}
	return sum
}
