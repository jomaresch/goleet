package top150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test122(t *testing.T) {
	tests := []struct {
		name           string
		prices         []int
		expectedProfit int
	}{
		{
			"sample 1",
			[]int{7, 1, 5, 3, 6, 4},
			7,
		},
		{
			"sample 2",
			[]int{1, 2, 3, 4, 5},
			4,
		},
		{
			"sample 3",
			[]int{7, 6, 4, 3, 1},
			0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedProfit, maxProfit2(test.prices))
		})
	}
}
