package main

// Given an array of strings nums containing n unique binary strings each of length n
// Return a binary string of length n that does not appear in nums
// If there are multiple answers, you may return any of them

func findDifferentBinaryString(nums []string) string {
	n := len(nums)
	res := make([]byte, n)
	for i, num := range nums {
		if num[i] == '0' {
			res[i] = '1'
		} else {
			res[i] = '0'
		}
	}
	return string(res)
}
