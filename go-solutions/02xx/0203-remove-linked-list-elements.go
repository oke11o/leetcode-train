package _2xx

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(node *ListNode, val int) *ListNode {
	result := &ListNode{Next: node}

	current := result
	for current.Next != nil {
		if current.Next.Val != val {
			current = current.Next
		} else {
			tmp := current.Next
			if tmp.Next != nil {
				current.Next = tmp.Next
			} else {
				current.Next = nil
			}
		}
	}
	return result.Next
}
