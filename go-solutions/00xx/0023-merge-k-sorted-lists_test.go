package _0xx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
https://leetcode.com/problems/merge-k-sorted-lists/
23. Merge k Sorted Lists
Hard
[#sorting, #linked_list]
*/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	to := len(lists)
	for to > 1 {
		for i := 0; i < to/2; i++ {
			lists[i] = merge2List(lists[i], lists[to-1-i])
		}
		to = to%2 + to/2
	}
	return lists[0]
}

func merge2List(a, b *ListNode) *ListNode {
	if a == nil {
		return b
	} else if b == nil {
		return a
	}
	dummy := &ListNode{}
	point := dummy
	for a != nil && b != nil {
		if a.Val < b.Val {
			point.Next = a
			a = a.Next
		} else {
			point.Next = b
			b = b.Next
		}
		point = point.Next
	}
	if a != nil {
		point.Next = a
	} else {
		point.Next = b
	}

	return dummy.Next
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				lists: [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			},
			want: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name: "",
			args: args{
				lists: [][]int{{1, 4, 5}, {1, 3, 4}},
			},
			want: []int{1, 1, 3, 4, 4, 5},
		},
		{
			name: "",
			args: args{
				lists: [][]int{},
			},
			want: []int{},
		},
		{
			name: "",
			args: args{
				lists: [][]int{{}},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pars := make([]*ListNode, len(tt.args.lists))
			for i, v := range tt.args.lists {
				pars[i] = BuildList(v)
			}
			assert.Equalf(t, tt.want, List2Slice(mergeKLists(pars)), "mergeKLists(%v)", tt.args.lists)
		})
	}
}
