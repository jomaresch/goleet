package top150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		name         string
		nums         []int
		val          int
		output       int
		expectedNums []int
	}{
		{
			"1",
			[]int{3, 2, 2, 3},
			3,
			2,
			[]int{2, 2},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.output, removeElement(test.nums, test.val))
			assert.Equal(t, test.expectedNums, test.nums[:len(test.expectedNums)])
		})
	}
}
