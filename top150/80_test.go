package top150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test80(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		output   int
		expected []int
	}{
		{
			"1",
			[]int{1, 2, 2, 3, 4, 4, 5},
			7,
			[]int{1, 2, 2, 3, 4, 4, 5},
		},
		{
			"one element",
			[]int{1},
			1,
			[]int{1},
		},
		{
			"one element twice",
			[]int{1, 1},
			2,
			[]int{1, 1},
		},
		{
			"two elements",
			[]int{1, 2},
			2,
			[]int{1, 2},
		},
		{
			"long",
			[]int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			7,
			[]int{0, 0, 1, 1, 2, 3, 3},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.output, removeDuplicates2(test.input))
			assert.Equal(t, test.expected, test.input[:len(test.expected)])
		})
	}
}
