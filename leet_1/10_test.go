package leet_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test10(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		result int
	}{
		{
			"1",
			[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			49,
		},
		{
			"2",
			[]int{1, 1},
			1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.result, maxArea(test.height))
		})
	}
}
