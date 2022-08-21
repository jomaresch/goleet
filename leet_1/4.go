package leet_1

import (
	"math"
)

// https://leetcode.com/problems/median-of-two-sorted-arrays/submissions/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	lenTotal := len1 + len2
	half := float64(lenTotal) / 2.0
	lower, upper := 0, 0
	if int(half) == int(math.Round(half)) {
		lower, upper = int(half), int(half)+1
	} else {
		lower, upper = int(half), int(half)
	}
	pointer1 := 0
	pointer2 := 0
	current := 0
	prev := 0
	for index := 0; index <= lower; index++ {
		if pointer1 < len1 {
			if pointer2 < len2 {
				if nums1[pointer1] < nums2[pointer2] {
					prev = current
					current = nums1[pointer1]
					pointer1++
				} else {
					prev = current
					current = nums2[pointer2]
					pointer2++
				}
				continue
			}
			prev = current
			current = nums1[pointer1]
			pointer1++
			continue
		}
		if pointer2 < len2 {
			prev = current
			current = nums2[pointer2]
			pointer2++
			continue
		}
	}
	if lower == upper {
		return float64(current)
	}
	return (float64(current) + float64(prev)) / 2
}
