package main

// Given an int array nums and a value val, remove all occurences of val in nums in place
// Return the new length of nums

// Solution: Keep track of the index of the last non-val element
// If the current element is not equal to val, then it is a non-val element
func removeElement(nums []int, val int) int {
	var k int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
