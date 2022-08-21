package leet_1

// https://leetcode.com/problems/add-two-numbers/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resStart := &ListNode{}
	resCurrent := resStart
	exceed := 0
	for {
		if l1 != nil {
			resCurrent.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			resCurrent.Val += l2.Val
			l2 = l2.Next
		}
		resCurrent.Val += exceed
		if resCurrent.Val > 9 {
			exceed = 1
			resCurrent.Val -= 10
		} else {
			exceed = 0
		}

		if l1 == nil && l2 == nil && exceed == 0 {
			return resStart
		}

		resCurrent.Next = &ListNode{}
		resCurrent = resCurrent.Next
	}
}
