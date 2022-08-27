package leet_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test5(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name:           "1",
			input:          "rtotr",
			expectedOutput: "rtotr",
		},
		{
			name:           "2",
			input:          "riror",
			expectedOutput: "rir",
		},
		{
			name:           "3",
			input:          "rbbor",
			expectedOutput: "bb",
		},
		{
			name:           "4",
			input:          "bb",
			expectedOutput: "bb",
		},
		{
			name:           "5",
			input:          "a",
			expectedOutput: "a",
		},
		{
			name:           "6",
			input:          "ab",
			expectedOutput: "b",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := longestPalindrome(test.input)
			assert.Equal(t, test.expectedOutput, res)
		})
	}
}
