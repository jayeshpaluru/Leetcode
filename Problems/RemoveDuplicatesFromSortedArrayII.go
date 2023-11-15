package main

// Given an int array nums sorted in a non-decreasing order
// Remove some duplicates in-place such that each unique element appears at most twice

// Solution: Keep track of the index of the last unique element
// If the current element is equal to the element two indices before it, then it is a duplicate
// Otherwise, it is a unique element
func removeDuplicatesII(nums []int) int {
	var k int
	for i := 0; i < len(nums); i++ {
		if k < 2 || nums[i] != nums[k-2] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
