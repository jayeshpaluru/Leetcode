package main

// The pair sum of a pair (a,b) is equal to a + b
// The maximum pair sum is the largest pair sum in an array
// Given an array nums of even length n, pair up the elements of the array into n/2 pairs
// Return the minimum pair sum of all the pairs

// Solution: Sort the array in ascending order
// The minimum pair sum is the sum of the smallest and largest elements in the array
// The smallest element is the first element in the array
// The largest element is the last element in the array

import "sort"

func minPairSum(nums []int) int {
	sort.Ints(nums)
	minPairSum := 0
	for i := 0; i < len(nums)/2; i++ {
		if nums[i]+nums[len(nums)-1-i] > minPairSum {
			minPairSum = nums[i] + nums[len(nums)-1-i]
		}
	}
	return minPairSum
}
