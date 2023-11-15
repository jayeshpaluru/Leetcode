package main

// Given two int arrays sorted in a non-decreasing order, and two integers m and n,
// Merge nums1 and nums2 into a single array sorted in a non-decreasing order.

// Solution: Merge nums2 into nums1, then sort nums1
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}

	for i := 0; i < m+n; i++ {
		for j := i + 1; j < m+n; j++ {
			if nums1[i] > nums1[j] {
				nums1[i], nums1[j] = nums1[j], nums1[i]
			}
		}
	}
}
