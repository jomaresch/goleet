package top150

func canJump(nums []int) bool {
	length := len(nums)
	if length == 1 {
		return true
	}
	for i := length - 1; i >= 0; i-- {
		if nums[i] != 0 {
			continue
		}
		if !checkOverJump(nums, i, i == length-1) {
			return false
		}
	}
	return true
}

func checkOverJump(nums []int, current int, isLast bool) bool {

	for index := current - 1; index >= 0; index-- {
		if isLast && index+nums[index] >= current {
			return true
		}
		if !isLast && index+nums[index] > current {
			return true
		}

	}
	return false
}
