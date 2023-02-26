package _0xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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

/*****************
****** TEST ******
*****************/
func Test_addTwoNumbers(t *testing.T) {
	var buildList = func(list []int) *ListNode {
		if len(list) == 0 {
			return nil
		}
		root := &ListNode{Val: list[0]}
		prev := root
		for i := 1; i < len(list); i++ {
			node := &ListNode{Val: list[i]}
			prev.Next = node
			prev = node
		}
		return root
	}

	tests := []struct {
		name string
		l1   []int
		l2   []int
		want []int
	}{
		{
			name: "",
			l1:   []int{2, 4, 3},
			l2:   []int{5, 6, 4},
			want: []int{7, 0, 8},
		},
		{
			name: "",
			l1:   []int{0},
			l2:   []int{0},
			want: []int{0},
		},
		{
			name: "",
			l1:   []int{9, 9, 9, 9, 9, 9, 9},
			l2:   []int{9, 9, 9, 9},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(buildList(tt.l1), buildList(tt.l2))
			require.Equal(t, buildList(tt.want), got)
		})
	}
}
