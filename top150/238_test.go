package top150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test238(t *testing.T) {
	tests := []struct {
		name           string
		input          []int
		expectedOutput []int
	}{
		{
			"1",
			[]int{1, 2, 3, 4},
			[]int{24, 12, 8, 6},
		}, {
			"2",
			[]int{1, 0, 3, 4},
			[]int{0, 12, 0, 0},
		}, {
			"3",
			[]int{1, 0, 3, 0},
			[]int{0, 0, 0, 0},
		}, {
			"4",
			[]int{0, 2, 3, 2},
			[]int{12, 0, 0, 0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedOutput, productExceptSelf(test.input))
		})
	}
}
