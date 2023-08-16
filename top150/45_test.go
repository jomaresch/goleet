package top150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_45(t *testing.T) {
	tests := []struct {
		name           string
		numsInput      []int
		expectedOutput int
	}{
		{
			name:           "1",
			numsInput:      []int{2, 3, 1, 1, 4},
			expectedOutput: 2,
		},
		{
			name:           "2",
			numsInput:      []int{2, 3, 0, 1, 4},
			expectedOutput: 2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedOutput, jump(test.numsInput))
		})
	}
}
