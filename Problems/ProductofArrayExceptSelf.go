package main

// Given an int array nums
// Return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i]
// Must run in O(n) time and without using division

// Solution: Use two arrays, left and right
// left[i] is the product of all the elements of nums before nums[i]
// right[i] is the product of all the elements of nums after nums[i]
// Then answer[i] = left[i] * right[i]
// We can use two variables to store the product of all the elements before and after nums[i] instead of two arrays
// Then we can compute the answer in one pass
func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	answer[0] = 1
	for i := 1; i < len(nums); i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}
	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] *= right
		right *= nums[i]
	}
	return answer
}
