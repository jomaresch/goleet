package leet_1

import (
	"math"
	"strconv"
)

// https://leetcode.com/problems/reverse-integer
func reverse(x int) int {
	isNegative := x < 0
	if isNegative {
		x *= -1
	}
	s := []rune(strconv.Itoa(x))
	l := len(s)

	for i := 0; i < (l / 2); i++ {
		s[i], s[l-i-1] = s[l-1-i], s[i]
	}
	n, _ := strconv.ParseInt(string(s), 10, 64)
	if n >= math.MaxInt32 {
		return 0
	}
	if isNegative {
		return int(n) * -1
	}
	return int(n)
}
