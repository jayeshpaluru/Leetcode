package main

// Given an int array nums and a value val, remove all occurences of val in nums in place
// Return the new length of nums
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
