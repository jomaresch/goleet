package top150

//https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/

func removeDuplicates2(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	setter := 2
	for index := 2; index < len(nums); index++ {
		if nums[index] == nums[setter-2] {
			continue
		}
		nums[setter] = nums[index]
		setter++
	}
	return setter
}
