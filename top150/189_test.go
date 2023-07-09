package top150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test189(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		rotations int
		expected  []int
	}{
		{
			"1",
			[]int{1, 2, 3, 4, 5, 6, 7},
			3,
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			"2",
			[]int{-1, -100, 3, 99},
			2,
			[]int{3, 99, -1, -100},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rotate(test.input, test.rotations)
			assert.Equal(t, test.expected, test.input)
		})
	}
}
