package leet_1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test12(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"III", 3},
		{"IV", 4},
		{"V", 5},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}
	for index, test := range tests {
		t.Run(fmt.Sprintf("test_%d", index), func(t *testing.T) {
			assert.Equal(t, test.output, romanToInt(test.input))
		})
	}
}
