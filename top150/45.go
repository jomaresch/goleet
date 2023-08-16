package top150

import (
	"math"
)

func jump(nums []int) int {
	// set variables
	jumps := 0
	maxDistance := 0
	currentPos := 0

	// loop through all items, except the last.
	for i := 0; i < len(nums)-1; i++ {
		// look for what is greater, current reach or new reachable.
		maxDistance = int(math.Max(float64(maxDistance), float64(nums[i]+i)))

		// if this iteration is the current step
		if currentPos == i {

			// increment jumps, set current position to maxDistance.
			jumps++
			currentPos = maxDistance
		}
	}

	// return total jumps
	return jumps
}
