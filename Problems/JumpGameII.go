package main

// Given a 0 indexed array of ints nums of length n, you initial position is nums[0]
// You can jump at most nums[i] steps at a time
// Return the minimum number of jumps to reach the last index
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
