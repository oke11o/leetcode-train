package _7xx

import (
	. "github.com/oke11o/leetcode-train/go-solutions/common"
)

// 1721. Swapping Nodes in a Linked List
// Medium
// Linked List, Two Pointers
// https://leetcode.com/problems/swapping-nodes-in-a-linked-list/
func swapNodes(head *ListNode, k int) *ListNode {
	result := &ListNode{Next: head}
	prevLeft := result
	finIndicator := head
	total := 0
	for finIndicator != nil {
		total++
		if total == k-1 {
			prevLeft = finIndicator
		}
		finIndicator = finIndicator.Next
	}

	prevRight := result
	finIndicator = head
	cnt := 0
	for finIndicator != nil {
		cnt++
		if cnt == total-k {
			prevRight = finIndicator
		}
		finIndicator = finIndicator.Next
	}
	if prevLeft == prevRight {
		return result.Next
	}

	if k > total/2 {
		switchNodes(prevRight, prevLeft)
	} else {
		switchNodes(prevLeft, prevRight)
	}

	return result.Next
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
}
