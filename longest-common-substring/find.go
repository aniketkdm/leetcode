package substring

// LongestCommonSubstring takes 2 strings and returns longest common sub-string between the two strings
// logic reference - https://www.youtube.com/watch?v=BysNXJHzCEs
func LongestCommonSubstring(s, r string) string {
	m := make([][]int, len(s))
	maxI := 0
	// maxJ := 0
	maxLen := 0

	for i := 0; i < len(r); i++ {
		m[i] = make([]int, len(s))
		for j := 0; j < len(s); j++ {
			if r[i] == s[j] {
				if i > 0 && j > 0 {
					m[i][j] = m[i-1][j-1] + 1
				} else {
					m[i][j] = 1
				}
				if m[i][j] > maxLen {
					maxLen = m[i][j]
					maxI = i
					// maxJ = j
				}
			}
		}
	}

	maxSubStr := make([]byte, maxLen)
	x := maxI
	for i := maxLen; i > 0; i-- {
		maxSubStr[i-1] = r[x]
		x--
	}

	return string(maxSubStr)
}
