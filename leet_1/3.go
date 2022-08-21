package leet_1

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	m := map[uint8]int{}
	longest := 0
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = i
			continue
		}

		if len(m) > longest {
			longest = len(m)
		}
		i = m[s[i]]
		m = map[uint8]int{}
	}
	if len(m) > longest {
		return len(m)
	}
	return longest
}

func lengthOfLongestSubstringSecond(s string) int {
	m := map[uint8]int{}
	longest := 0
	second := 0
	for i := 0; i < len(s); i++ {
		foundIndex, ok := m[s[i]]
		if !ok {
			m[s[i]] = i
			continue
		}

		if len(m) > longest {
			longest = len(m)
		}

		for s2 := second; s2 <= foundIndex; s2++ {
			delete(m, s[s2])
			second = s2 + 1
		}

		m[s[i]] = i
	}
	if len(m) > longest {
		return len(m)
	}
	return longest
}
