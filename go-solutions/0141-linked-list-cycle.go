package go_solutions

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast := head
	slow := head
	for {
		if slow.Next == nil {
			return false
		}
		slow = slow.Next

		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}
		fast = fast.Next.Next

		if slow == fast {
			return true
		}

	}
}
