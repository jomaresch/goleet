package top150

func productExceptSelf(nums []int) []int {
	totalProduct := nums[0]
	zeroCount := 0
	startIndex := 1

	if totalProduct == 0 {
		totalProduct = nums[1]
		zeroCount = 1
		startIndex = 2
	}

	for index := startIndex; index < len(nums); index++ {
		if nums[index] == 0 {
			zeroCount++
			continue
		}
		totalProduct *= nums[index]
	}

	result := make([]int, len(nums))
	if zeroCount >= 2 {
		return result
	}
	for index, num := range nums {
		if zeroCount > 0 {
			if num == 0 {
				result[index] = totalProduct
			}
			continue
		}
		result[index] = totalProduct / num
	}
	return result
}
