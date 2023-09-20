package _0xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

/*
*
92. Reverse Linked List II
Medium
Topics: Linked List
Companies
Given the head of a singly linked list and two integers left and right where left <= right, reverse the nodes of the
list from position left to position right, and return the reversed list
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var reverse = func(h *ListNode) *ListNode {
		var prev *ListNode
		for h != nil {
			next := h.Next
			h.Next = prev
			prev = h
			h = next
		}
		return prev
	}
	_ = reverse

	result := head
	var prev *ListNode
	var bPrev *ListNode
	var bCur *ListNode
	i := 0
	for head != nil {
		next := head.Next
		i++
		if i == left {
			bPrev = prev
			bCur = head
		} else if i == right {
			if bPrev != nil {
				bPrev.Next = head
			} else {
				result = head
			}
			bCur.Next = next
		}

		if i > left && i <= right {
			head.Next = prev
		}
		prev = head
		head = next
	}

	return result
}

func Test_reverseBetween(t *testing.T) {
	tests := []struct {
		name  string
		head  []int
		left  int
		right int
		want  []int
	}{
		{
			name:  "",
			head:  []int{1, 2, 3, 4, 5},
			left:  1,
			right: 5,
			want:  []int{5, 4, 3, 2, 1},
		},
		{
			name:  "",
			head:  []int{1, 2, 3, 4, 5},
			left:  1,
			right: 4,
			want:  []int{4, 3, 2, 1, 5},
		},
		{
			name:  "",
			head:  []int{1, 2, 3, 4, 5},
			left:  2,
			right: 4,
			want:  []int{1, 4, 3, 2, 5},
		},
		{
			name:  "",
			head:  []int{5},
			left:  1,
			right: 1,
			want:  []int{5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := BuildList(tt.head)
			got := reverseBetween(list, tt.left, tt.right)
			got2 := List2Slice(got)
			require.Equal(t, tt.want, got2)
		})
	}
}
