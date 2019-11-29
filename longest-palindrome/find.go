package palindrome

import (
	substring "github.com/aniketkdm/leetcode/longest-common-substring"
)

func longestPalindrome(s string) string {
	rev := reverse(s)
	commSubStr := substring.LongestCommonSubstring(s, rev)

	return ""
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
