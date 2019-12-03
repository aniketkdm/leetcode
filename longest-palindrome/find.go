package palindrome

func longestPalindrome(s string) string {
	m := fillMatrix(s)

	return longestPalindromeFromSubstrings(m, s)
}

func longestPalindromeFromSubstrings(m [][]int, s string) string {
	var longestPalindrome string

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] > 0 && m[i][j] > len(longestPalindrome) {
				subStr := getSubstring(m, s, i, j)
				if isPalindrome(subStr, 0, len(subStr)-1) && len(subStr) > len(longestPalindrome) {
					longestPalindrome = subStr
				}
			}
		}
	}

	return longestPalindrome
}

func getSubstring(m [][]int, s string, i, j int) string {
	subStr := make([]byte, m[i][j])
	for x := m[i][j]; x > 0; x-- {
		subStr[x-1] = s[j]
		j--
	}
	return string(subStr)
}

func isPalindrome(s string, start, end int) bool {
	if start == end || end < start {
		return true
	}
	if s[start] == s[end] {
		return isPalindrome(s, start+1, end-1)
	}
	return false
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func fillMatrix(s string) [][]int {
	m := make([][]int, len(s))
	r := reverse(s)

	for i := 0; i < len(r); i++ {
		m[i] = make([]int, len(s))
		for j := 0; j < len(s); j++ {
			if r[i] == s[j] {
				if i > 0 && j > 0 {
					m[i][j] = m[i-1][j-1] + 1
				} else {
					m[i][j] = 1
				}
			}
		}
	}

	return m
}
