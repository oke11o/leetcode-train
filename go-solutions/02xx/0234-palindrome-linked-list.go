package _2xx

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	middle := middleNode(head)
	reverse := reverseList(middle)

	for reverse != nil {
		if reverse.Val != head.Val {
			return false
		}
		reverse = reverse.Next
		head = head.Next
	}

	return true
}
func middleNode(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		} else {
			return slow
		}
		slow = slow.Next
	}
	return slow
}
