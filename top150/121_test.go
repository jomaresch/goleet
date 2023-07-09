package top150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test121(t *testing.T) {
	tests := []struct {
		name           string
		prices         []int
		expectedProfit int
	}{
		{
			"1",
			[]int{7, 1, 5, 3, 6, 4},
			5,
		}, {
			"negative",
			[]int{7, 6, 4, 3, 1},
			0,
		}, {
			"odd",
			[]int{1, 4, 2},
			3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedProfit, maxProfit(test.prices))
		})
	}
}
