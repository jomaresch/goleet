package leet_1

func maxArea(height []int) int {
	maxA := 0
	right := len(height) - 1
	left := 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		if maxA < area {
			maxA = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxA
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
