package _0xx

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersRec(l1, l2, false)
}

func addTwoNumbersRec(l1 *ListNode, l2 *ListNode, add bool) *ListNode {
	if l1 == nil && l2 == nil {
		if !add {
			return nil
		}
		return &ListNode{Val: 1}
	}
	l1Val := 0
	var l1Child *ListNode
	if l1 != nil {
		l1Val = l1.Val
		l1Child = l1.Next
	}
	l2Val := 0
	var l2Child *ListNode
	if l2 != nil {
		l2Val = l2.Val
		l2Child = l2.Next
	}

	val := l1Val + l2Val
	if add {
		val++
	}
	if val > 9 {
		val = val % 10
		add = true
	} else {
		add = false
	}
	node := &ListNode{
		Val:  val,
		Next: addTwoNumbersRec(l1Child, l2Child, add),
	}

	return node
}
