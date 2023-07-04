package leet_1

func longestCommonPrefix(strs []string) string {
	sLength := len(strs)
	result := ""
	for charIndex := range strs[0] {
		for wordIndex := 1; wordIndex < sLength; wordIndex++ {
			if charIndex >= len(strs[wordIndex]) || strs[0][charIndex] != strs[wordIndex][charIndex] {
				return result
			}
		}
		result += string(strs[0][charIndex])
	}
	return result
}
