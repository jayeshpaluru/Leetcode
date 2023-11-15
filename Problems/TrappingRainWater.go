package main

// Given n non-negative integers representing an elevation map where the width of each bar is 1
// Compute how much water it can trap after raining.

// Solution: Dynamic Programming
// For each bar, the water it can trap is the minimum of the highest bar on its left and right minus its height.
// So we can use two arrays to store the highest bar on its left and right.
// Then we can compute the water it can trap for each bar.
func trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	left, right := make([]int, n), make([]int, n)
	left[0], right[n-1] = height[0], height[n-1]
	for i := 1; i < n; i++ {
		left[i] = max(left[i-1], height[i])
		right[n-i-1] = max(right[n-i], height[n-i-1])
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += min(left[i], right[i]) - height[i]
	}
	return sum
}
