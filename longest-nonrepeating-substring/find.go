package substring

import "strings"

func lengthOfLongestSubstring(s string) int {
	// emptyStruct := struct{}{}
	// instead of map, an array of length 128 (ascii values for chars) is more space efficient and probably faster
	m := make(map[string]int)
	maxLen := 0
	currLen := 0

	strArr := strings.Split(s, "")

	for i := 0; i < len(strArr); i++ {
		if n, ok := m[strArr[i]]; !ok {
			m[strArr[i]] = i
			currLen = currLen + 1
		} else {
			if currLen > maxLen {
				maxLen = currLen
			}
			// i = i - currLen + 1
			i = n + 1
			currLen = 1
			m = map[string]int{strArr[i]: i}
		}
	}
	if currLen > maxLen {
		maxLen = currLen
	}
	return maxLen
}
