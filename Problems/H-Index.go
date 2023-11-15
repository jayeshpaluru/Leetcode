package main

// Given an int array citations where citations[i] is the number of citations a researcher received for their ith paper
// Return the researcher's h-index

// Solution: Sort the array in descending order
// The h-index is the largest number h such that there are at least h papers with h citations
// If the current citation is greater than or equal to the current index, then it is a valid h-index
// Otherwise, it is not a valid h-index
func hIndex(citations []int) int {
	for i := 0; i < len(citations); i++ {
		for j := i; j > 0 && citations[j] > citations[j-1]; j-- {
			citations[j], citations[j-1] = citations[j-1], citations[j]
		}
	}
	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] >= i+1 {
			return i + 1
		}
	}
	return 0
}
