package _0xx

import "github.com/oke11o/leetcode-train/go-solutions"

/**
 * 0021. Merge Two Sorted Lists
 * Easy
 * Linked List, Recursion
 * https://leetcode.com/problems/merge-two-sorted-lists/
 */
func mergeTwoLists(l1 *go_solutions.ListNode, l2 *go_solutions.ListNode) *go_solutions.ListNode {
	result := &go_solutions.ListNode{}
	current := result

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			current = current.Next
			l1 = l1.Next
		} else {
			current.Next = l2
			current = current.Next
			l2 = l2.Next
		}
	}
	for l1 != nil {
		current.Next = l1
		current = current.Next
		l1 = l1.Next
	}
	for l2 != nil {
		current.Next = l2
		current = current.Next
		l2 = l2.Next
	}

	return result.Next
}
