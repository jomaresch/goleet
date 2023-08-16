package top150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_274(t *testing.T) {
	tests := []struct {
		name           string
		input          []int
		expectedOutput int
	}{
		{
			"1",
			[]int{3, 0, 6, 1, 5},
			3,
		}, {
			"2",
			[]int{1, 3, 1},
			1,
		},
		{
			"3",
			[]int{11, 15},
			2,
		},
		{
			"4",
			[]int{4, 4, 0, 0},
			2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedOutput, hIndex(test.input))
		})
	}
}
