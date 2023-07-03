package leet_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test11(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output string
	}{
		{
			"1",
			3,
			"III",
		},
		{
			"2",
			58,
			"LVIII",
		},
		{
			"1",
			1994,
			"MCMXCIV",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.output, intToRoman(test.input))
		})
	}
}
