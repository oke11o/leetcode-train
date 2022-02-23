package _0xx

// 0019. Remove Nth Node From End of List
// Medium
// Linked List, Two Pointers
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result := &ListNode{Next: head}
	fast := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	slow := head
	prev := result
	for fast != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next
	}
	if prev != nil && prev.Next != nil {
		prev.Next = prev.Next.Next
	}

	return result.Next
}
