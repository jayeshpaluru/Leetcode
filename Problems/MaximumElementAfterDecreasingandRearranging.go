package main

// Given an array of integers arr
// Perform some operations (possibly none) on arr so that it satisfies these conditions:
// 1. The value of the first element in arr must be 1.
// 2. The absolute difference between any 2 adjacent elements must be less than or equal to 1.
// There are 2 types of operations that you can perform any number of times:
// 1. Decrease the value of any element of arr to a smaller positive integer.
// 2. Rearrange the elements of arr to be in any order.
// Return the maximum possible value of an element in arr after performing the operations to satisfy the conditions.
import "sort"

// This solution works by sorting the array and then iterating through it.
// For each element, we check if the difference between the current element and the previous element is greater than 1.
// If it is, then we set the current element to the previous element plus 1.
// At the end, we return the last element in the array.
func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	arr[0] = 1
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] > 1 {
			arr[i] = arr[i-1] + 1
		}
	}
	return arr[len(arr)-1]
}
