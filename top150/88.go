package top150

// https://leetcode.com/problems/merge-sorted-array

func merge(nums1 []int, m int, nums2 []int, n int) {
	for p := m + n - 1; n > 0; p-- {
		if m > 0 && nums1[m-1] > nums2[n-1] {
			nums1[p] = nums1[m-1]
			m--
		} else {
			nums1[p] = nums2[n-1]
			n--
		}
	}
}
