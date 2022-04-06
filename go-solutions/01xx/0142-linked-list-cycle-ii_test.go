package _1xx

import (
	"reflect"
	"testing"
)

/**
https://leetcode.com/problems/linked-list-cycle-ii/
142. Linked List Cycle II
Medium
*/
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast := head
	slow := head
	for {
		if slow.Next == nil {
			return nil
		}
		slow = slow.Next

		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	// here is finding
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
