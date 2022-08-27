package leet_1

// https://leetcode.com/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return string(s[0])
	}
	longest := ""
	for i := 1; i < len(s); i++ {
		// Check odd palindromes
		right, left := i, i
		for right+1 < len(s) && left > 0 && s[left-1] == s[right+1] {
			right++
			left--
		}
		if right+1-left > len(longest) {
			longest = s[left : right+1]
		}

		// Check even palindromes
		right, left = i-1, i
		if s[left] != s[right] {
			continue
		}
		for right+1 < len(s) && left > 0 && s[left-1] == s[right+1] {
			right++
			left--
		}
		if right+1-left > len(longest) {
			longest = s[left : right+1]
		}
	}
	return longest
}
