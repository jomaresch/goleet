package leet_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		name           string
		inputNumbers   []int
		inputTarget    int
		expectedOutput []int
	}{
		{
			name:           "1",
			inputNumbers:   []int{1, 4, 7, 11, 88, 122},
			inputTarget:    92,
			expectedOutput: []int{1, 4},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := twoSum(test.inputNumbers, test.inputTarget)
			assert.Equal(t, test.expectedOutput, result)
		})
	}
}
