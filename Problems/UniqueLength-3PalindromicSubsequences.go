package main

// Given a string s
// Return the number of unique palindromes of length three that are a subsequence of s.

// Helper function that takes a string and returns a map where each key is a unique character.
func unique(s string) map[rune]bool {
	set := map[rune]bool{}
	for _, c := range s {
		set[c] = true
	}
	return set
}

// Works by iterating through the string and keeping track of the leftmost and rightmost index of each character.
// Then, for each character, we check if there are any characters between the leftmost and rightmost index of that character.
// If there are, then we know that we can form a palindrome of length three with that character.
func countPalindromicSubsequence(s string) int {
	left, right := [26]int{}, [26]int{}
	for i := range left {
		left[i] = len(s)
	}

	for i, c := range s {
		right[c-'a'] = i
	}

	for i, c := range s {
		if i < left[c-'a'] {
			left[c-'a'] = i
		}
	}

	count := 0
	for c := 0; c < 26; c++ {
		if left[c] < right[c] {
			count += len(unique(s[left[c]+1 : right[c]]))
		}
	}

	return count
}
