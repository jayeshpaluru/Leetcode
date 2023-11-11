package main

// Given an int array nums of size n, return the majority elemenet
// The majority element is the element that appears more than ⌊n/2⌋ times

func majorityElement(nums []int) int {
	var candidate, i int
	for _, num := range nums {
		if i == 0 {
			candidate = num
		}
		if num == candidate {
			i++
		} else {
			i--
		}
	}
	return candidate
}
