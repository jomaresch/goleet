package top150

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test88(t *testing.T) {
	tests := []struct {
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			[]int{1, 2, 3, 0, 0, 0},
			3,
			[]int{2, 5, 6},
			3,
			[]int{1, 2, 2, 3, 5, 6},
		},
		{
			[]int{1},
			1,
			[]int{},
			0,
			[]int{1},
		},
		{
			[]int{0},
			0,
			[]int{1},
			1,
			[]int{1},
		},
	}

	for index, test := range tests {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			merge(test.nums1, test.m, test.nums2, test.n)
			assert.Equal(t, test.expected, test.nums1)
		})
	}
}
