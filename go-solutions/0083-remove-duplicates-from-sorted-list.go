package go_solutions

func deleteDuplicates(head *ListNode) *ListNode {
	result := &ListNode{Val: -99999}
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
