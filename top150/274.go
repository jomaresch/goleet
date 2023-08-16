package top150

import (
	"sort"
)

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	count := 0
	for _, c := range citations {
		if c > count {
			count++
		}
	}

	return count
}
