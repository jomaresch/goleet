package leet_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test6(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		input2         int
		expectedOutput string
	}{
		{
			name:           "1",
			input:          "PAYPALISHIRING",
			input2:         3,
			expectedOutput: "PAHNAPLSIIGYIR",
		},
		{
			name:           "2",
			input:          "PAYPALISHIRING",
			input2:         1,
			expectedOutput: "PAYPALISHIRING",
		},
		{
			name:           "3",
			input:          "PAYPALISHIRING",
			input2:         4,
			expectedOutput: "PINALSIGYAHRPI",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := convert(test.input, test.input2)
			assert.Equal(t, test.expectedOutput, res)
		})
	}
}
