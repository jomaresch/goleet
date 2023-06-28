package leet_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test8(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output bool
	}{
		{
			name:   "success uneven",
			input:  1234321,
			output: true,
		},
		{
			name:   "success even",
			input:  123321,
			output: true,
		},
		{
			name:   "fail",
			input:  1234,
			output: false,
		},
		{
			name:   "negative",
			input:  -1,
			output: false,
		},
		{
			name:   "one digit",
			input:  8,
			output: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.output, isPalindrome(test.input))
		})
	}
}
