package leet_1

import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	reversedX := 0
	tmp := x
	for tmp > 0 {
		reversedX = reversedX*10 + tmp%10
		tmp /= 10
	}
	return reversedX == x
}

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	xString := strconv.Itoa(x)
	xStringLength := len(xString)
	for i := 0; i <= xStringLength/2; i++ {
		if xString[i] != xString[xStringLength-1-i] {
			return false
		}
	}

	return true
}
