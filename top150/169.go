package top150

//https://leetcode.com/problems/majority-element/

func majorityElement(nums []int) int {
	currentElement, voteCount := nums[0], 1
	for index := 1; index < len(nums); index++ {
		if voteCount == 0 {
			currentElement, voteCount = nums[index], 1
			continue
		}

		if nums[index] == currentElement {
			voteCount++
		} else {
			voteCount--
		}
	}
	return currentElement
}
