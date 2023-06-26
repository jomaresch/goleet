package leet_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test7(t *testing.T) {
	tests := []struct {
		name           string
		input          int
		expectedOutput int
	}{
		{
			name:           "1",
			input:          1234,
			expectedOutput: 4321,
		},
		{
			name:           "negative",
			input:          -1234,
			expectedOutput: -4321,
		},
		{
			name:           "too large",
			input:          1534236469,
			expectedOutput: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedOutput, reverse(test.input))
		})
	}
}
