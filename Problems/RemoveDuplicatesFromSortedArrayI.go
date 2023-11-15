package main

// Given an int array nums sorted in a non-decreasing order
// Remove the duplicates in-place such that each unique element appears once

// Solution: Keep track of the index of the last unique element
// If the current element is not equal to the last unique element, then it is a unique element
func removeDuplicates(nums []int) int {
	var k int
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[k] {
			k++
			nums[k] = nums[i]
		}
	}
	return k + 1
}
