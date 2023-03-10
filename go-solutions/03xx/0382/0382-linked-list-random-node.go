package _382

import (
	"math/rand"

	. "github.com/oke11o/leetcode-train/go-solutions/common"
)

/*
https://leetcode.com/problems/linked-list-random-node/
382. Linked List Random Node
Medium
*/

type Solution struct {
	head   *ListNode
	length int
}

func Constructor(head *ListNode) Solution {
	length := 0
	h := head
	for h != nil {
		length++
		h = h.Next
	}
	return Solution{head: head, length: length}
}

func (this *Solution) GetRandom() int {
	p := rand.Intn(this.length)
	h := this.head
	for p > 0 {
		p--
		h = h.Next
	}

	return h.Val
}
