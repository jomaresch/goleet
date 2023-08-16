package top150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanJump(t *testing.T) {
	tests := []struct {
		name           string
		numsInput      []int
		expectedResult bool
	}{
		{
			name:           "success",
			numsInput:      []int{2, 3, 1, 1, 4},
			expectedResult: true,
		}, {
			name:           "success 2",
			numsInput:      []int{2, 5, 0, 0},
			expectedResult: true,
		}, {
			name:           "success 3",
			numsInput:      []int{2, 0, 0},
			expectedResult: true,
		}, {
			name:           "success 4",
			numsInput:      []int{2, 0},
			expectedResult: true,
		},
		{
			name:           "fail",
			numsInput:      []int{3, 2, 1, 0, 4},
			expectedResult: false,
		},
		{
			name:           "edge",
			numsInput:      []int{0},
			expectedResult: true,
		},
		{
			name:           "edge 2",
			numsInput:      []int{1},
			expectedResult: true,
		},
		{
			name:           "fail 2",
			numsInput:      []int{0, 1},
			expectedResult: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedResult, canJump(test.numsInput))
		})
	}
}
