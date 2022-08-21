package leet_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test4(t *testing.T) {
	tests := []struct {
		name           string
		input1         []int
		input2         []int
		expectedOutput float64
	}{
		{
			name:           "1",
			input1:         []int{1, 3, 5, 7, 9},
			input2:         []int{1, 3, 5, 7, 9},
			expectedOutput: 5,
		},
		{
			name:           "2",
			input1:         []int{1, 7, 9},
			input2:         []int{1, 3, 5, 7, 9},
			expectedOutput: 6,
		},
		{
			name:           "3",
			input1:         []int{1, 3},
			input2:         []int{2},
			expectedOutput: 2,
		},
		{
			name:           "4",
			input1:         []int{1, 4},
			input2:         []int{2, 3},
			expectedOutput: 2.5,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := findMedianSortedArrays(test.input1, test.input2)
			assert.Equal(t, test.expectedOutput, res)
		})
	}
}
