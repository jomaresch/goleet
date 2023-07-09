package top150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test169(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			"one element",
			[]int{3},
			3,
		},
		{
			"three elements",
			[]int{3, 9, 9},
			9,
		},
		{
			"multiple elements",
			[]int{2, 2, 1, 1, 1, 2, 2},
			2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, majorityElement(test.input))
		})
	}
}
