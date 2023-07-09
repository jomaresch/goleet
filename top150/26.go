package top150

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {
	setter := 1
	for index := 1; index < len(nums); index++ {
		if nums[index] == nums[setter-1] {
			continue
		}
		nums[setter] = nums[index]
		setter++
	}
	return setter
}
