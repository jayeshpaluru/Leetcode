package main

// Given an int array nums sorted in a non-decreasing order
// Remove some duplicates in-place such that each unique element appears at most twice

func removeDuplicates(nums []int) int {
	var k int
	for i := 0; i < len(nums); i++ {
		if k < 2 || nums[i] != nums[k-2] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
