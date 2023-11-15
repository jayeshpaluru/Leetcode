package main

// Given an int array nums
// Rotate the array to the right by k steps, where k is non-negative
// I'm going to make a helper function called reverse that reverses a slice of ints
func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// Solution: Reverse the entire array, then reverse the first k elements, then reverse the rest of the elements
func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}
