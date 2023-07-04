package leet_1

import (
	"strings"
)

func romanToInt(s string) int {
	mappingIndex := len(romanMapping) - 1
	result := 0
	for len(s) > 0 {
		if strings.HasPrefix(s, romanMapping[mappingIndex].val) {
			result += romanMapping[mappingIndex].num
			s = strings.TrimPrefix(s, romanMapping[mappingIndex].val)
		} else {
			mappingIndex--
		}
	}
	return result
}
