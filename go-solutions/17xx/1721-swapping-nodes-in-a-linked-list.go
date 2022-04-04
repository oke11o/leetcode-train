package _7xx

import (
	. "github.com/oke11o/leetcode-train/go-solutions/common"
)

// 1721. Swapping Nodes in a Linked List
// Medium
// Linked List, Two Pointers
// https://leetcode.com/problems/swapping-nodes-in-a-linked-list/
func swapNodes(head *ListNode, k int) *ListNode {
	// find length. (As backside effect, search prev for k from left)
	var front *ListNode
	var end *ListNode
	current := head
	length := 0
	// set the front node and end node in single pass
	for current != nil {
		length++        // так как у нас 1 indexed list
		if end != nil { // мы установим end, когда дойдем до k. А потом просто буду двигать end так же как current
			end = end.Next
		}
		// check if we have reached kth node
		if length == k {
			front = current
			end = head
		}
		current = current.Next
	}
	front.Val, end.Val = end.Val, front.Val

	return head
}

func switchNodes(prevLeft *ListNode, prevRight *ListNode) {
	right := prevRight.Next
	left := prevLeft.Next

	tmpRightNext := right.Next

	// 1. Разорвать левую prev связь. Установивь на правую. Но! На правую ссылаются 2. prevRight & prevLeft
	prevLeft.Next = right
	// 2. Установить новую prevLeft связь. Разорвать лишнюю prevRight связь
	prevRight.Next = left
	// 3. Устновить новую nextRight связь.
	right.Next = left.Next
	// 4. Set new nextLeft connect
	left.Next = tmpRightNext

	// Можно проще. По сути, мы знаем, куда и что вставлять.
	// Удаляем текущую ноду.
	// Вставляем в нужное место нужую ноду.
	// Тогда не надо знать, какая из них левая, а какая правая

	// А еще лучше, можно просто Val менять. А не указатели

}
