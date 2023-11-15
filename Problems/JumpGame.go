package main

// Given an int array nums, you are initially positioned at the array's first index
// Each element in the array represents your maximum jump length at that position, return true if you can reach the last index

// Solution: Use one variable, maxJump
// maxJump is the maximum index you can reach from the current index
// If the current index is greater than maxJump, then you cannot reach the current index
// If the current index plus the maximum jump length is greater than maxJump, then you can reach the current index
func canJump(nums []int) bool {
	var maxJump int
	for i := 0; i < len(nums); i++ {
		if i > maxJump {
			return false
		}
		if i+nums[i] > maxJump {
			maxJump = i + nums[i]
		}
	}
	return true
}
