package main

// Given a string s
// Return the number of unique palindromes of length three that are a subsequence of s.
import "strings"

func countPalindromicSubsequence(s string) int {
	var count int
	for i := 0; i < 26; i++ {
		first := strings.IndexByte(s, byte(i+'a'))
		last := strings.LastIndexByte(s, byte(i+'a'))
		if first != -1 && last != -1 && first != last {
			for j := 0; j < 26; j++ {
				if strings.IndexByte(s[first+1:last], byte(j+'a')) != -1 {
					count++
					break
				}
			}
		}
	}
	return count
}
