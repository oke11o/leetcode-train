package _0xx

import (
	go_solutions "github.com/oke11o/leetcode-train/go-solutions/01xx"
)

func deleteDuplicates(head *go_solutions.ListNode) *go_solutions.ListNode {
	result := &go_solutions.ListNode{Val: -99999}
	result.Next = head

	current := result
	for current.Next != nil {
		if current.Next.Val != current.Val {
			current = current.Next
		} else {
			if current.Next.Next != nil {
				current.Next = current.Next.Next
			} else {
				current.Next = nil
			}
		}
	}

	return result.Next
}
