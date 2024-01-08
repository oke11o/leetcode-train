package _8xx

/*
876. Middle of the Linked List
https://leetcode.com/problems/middle-of-the-linked-list/
*/
func middleNode(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
