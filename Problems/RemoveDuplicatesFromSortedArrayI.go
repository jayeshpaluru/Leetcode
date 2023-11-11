package main

// Given an int array nums sorted in a non-decreasing order
// Remove the duplicates in-place such that each unique element appears once
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
