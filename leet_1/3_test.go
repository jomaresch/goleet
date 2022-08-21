package leet_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test3(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput int
	}{
		{
			name:           "1",
			input:          "abcdeehijklmnopihy",
			expectedOutput: 10,
		}, {
			name:           "2",
			input:          "pwwkew",
			expectedOutput: 3,
		}, {
			name:           "3",
			input:          "bpfbhmipx",
			expectedOutput: 7,
		}, {
			name:           "4",
			input:          "anviaj",
			expectedOutput: 5,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := lengthOfLongestSubstring(test.input)
			assert.Equal(t, test.expectedOutput, res)

			res2 := lengthOfLongestSubstringSecond(test.input)
			assert.Equal(t, test.expectedOutput, res2)
		})
	}
}
