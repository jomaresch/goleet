package leet_1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test13(t *testing.T) {
	tests := []struct {
		input  []string
		output string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"ilower", "olow", "flight"}, ""},
	}
	for index, test := range tests {
		t.Run(fmt.Sprintf("test_%d", index), func(t *testing.T) {
			assert.Equal(t, test.output, longestCommonPrefix(test.input))
		})
	}
}
