package main

// Given a 0 indexed array of ints nums of length n, you initial position is nums[0]
// You can jump at most nums[i] steps at a time
// Return the minimum number of jumps to reach the last index

// Solution: Use two variables, maxJump and end
// maxJump is the maximum index you can reach from the current index
// end is the maximum index you can reach from the current jump
// If the current index is equal to end, then you have reached the end of the current jump
// So you must jump again
// The number of jumps is equal to the number of times you reach the end of a jump
func jump(nums []int) int {
	var maxJump, jumps, end int
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > maxJump {
			maxJump = i + nums[i]
		}
		if i == end {
			jumps++
			end = maxJump
		}
	}
	return jumps
}
