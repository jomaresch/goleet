package leet_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newLinkedList(items ...int) *ListNode {
	start := &ListNode{}
	node := start
	for index, i := range items {
		node.Val = i
		if index+1 == len(items) {
			break
		}
		node.Next = &ListNode{}
		node = node.Next
	}
	return start
}

func Test_2(t *testing.T) {
	tests := []struct {
		name           string
		inputA         *ListNode
		inputB         *ListNode
		outputExpected *ListNode
	}{
		{
			name:           "1",
			inputA:         newLinkedList(9, 8, 7, 6),
			inputB:         newLinkedList(1),
			outputExpected: newLinkedList(0, 9, 7, 6),
		},
		{
			name:           "2",
			inputA:         newLinkedList(),
			inputB:         newLinkedList(),
			outputExpected: newLinkedList(),
		},
		{
			name:           "3",
			inputA:         newLinkedList(9, 9, 9, 9),
			inputB:         newLinkedList(1),
			outputExpected: newLinkedList(0, 0, 0, 0, 1),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := addTwoNumbers(test.inputA, test.inputB)
			assert.Equal(t, test.outputExpected, result)
		})
	}
}
