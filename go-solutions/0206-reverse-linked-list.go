package go_solutions

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(node *ListNode) *ListNode {
	var result *ListNode
	for node != nil {
		next := node.Next
		node.Next = result
		result = node
		node = next
	}
	return result
}
