package top150

// https://leetcode.com/problems/remove-element/

func removeElement(nums []int, val int) int {
	setIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[setIndex] = nums[i]
			setIndex++
		}
	}
	copy(nums, nums[:setIndex])
	return setIndex
}
