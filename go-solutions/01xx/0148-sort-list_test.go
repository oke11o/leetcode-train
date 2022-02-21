package _1xx

import (
	"reflect"
	"testing"

	"github.com/oke11o/leetcode-train/go-solutions/common"
)

/**
148. Sort List
Medium
https://leetcode.com/problems/sort-list/
TODO: not working
*/
func sortList(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var merge = func(a, b *common.ListNode) *common.ListNode {
		dummy := &common.ListNode{}
		tail := dummy
		for a != nil && b != nil {
			if a.Val < b.Val {
				tail.Next = a
				tail = tail.Next
				a = a.Next
			} else {
				tail.Next = b
				tail = tail.Next
				b = b.Next
			}
		}

		if a != nil {
			tail.Next = a
		} else {
			tail.Next = b
		}
		return dummy.Next
	}
	var getMid = func(a *common.ListNode) *common.ListNode {
		fast := head
		slow := head
		for fast != nil {
			fast = fast.Next
			if fast != nil {
				fast = fast.Next
			} else {
				return slow
			}
			slow = slow.Next
		}
		return slow
	}
	mid := getMid(head)
	left := sortList(head)
	right := sortList(mid)
	return merge(left, right)
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_sortList(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				head: []int{4, 2, 1, 3},
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "",
			args: args{
				head: []int{-1, 5, 3, 4, 0},
			},
			want: []int{-1, 0, 3, 4, 5},
		},
		{
			name: "",
			args: args{
				head: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortList(common.BuildList(tt.args.head))
			if !reflect.DeepEqual(common.List2Slice(got), tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
