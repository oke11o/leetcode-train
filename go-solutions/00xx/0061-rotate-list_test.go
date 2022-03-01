package _0xx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
https://leetcode.com/problems/rotate-list/
61. Rotate List
Medium
*/
/**
- Запоминаю первый элемент
- Иду до
 - либо до k
 - либо до len
   (что меньше)
- если k меньше. То запоминаю его. У предыдущего ставилю child = nil. И иду до конца. У конца. У конца ствалю child=head

- если len меньше. Вычисляю k = k%len. И повтораю выше алгоритм
И это работает для сдвига назад. Нам же надо сдвинуть вперед



*/
func rotateRight(head *ListNode, k int) *ListNode {
	var getLenght = func(head *ListNode) int {
		result := 0
		for head != nil {
			head = head.Next
			result++
		}
		return result
	}

	if head == nil || head.Next == nil {
		return head
	}
	length := getLenght(head)
	k = k % length
	if k == 0 {
		return head
	}
	k = length - k // ВАЖНО!!! Голова должна указывать на k элемент с конца.

	cur := head
	newHead := cur
	var prev *ListNode
	for cur != nil {
		prev = cur
		cur = cur.Next
		k--
		if 0 == k {
			newHead = cur
			prev.Next = nil
		}
		if cur == nil {
			prev.Next = head
		}
	}
	return newHead
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_rotateRight(t *testing.T) {
	type args struct {
		head []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				head: []int{1, 2},
				k:    1,
			},
			want: []int{2, 1},
		},
		{
			name: "",
			args: args{
				head: []int{1, 2},
				k:    2,
			},
			want: []int{1, 2},
		},
		{
			name: "",
			args: args{
				head: []int{0, 1, 2},
				k:    1,
			},
			want: []int{2, 0, 1},
		},
		{
			name: "",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				k:    2,
			},
			want: []int{4, 5, 1, 2, 3},
		},
		{
			name: "",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				k:    7,
			},
			want: []int{4, 5, 1, 2, 3},
		},
		{
			name: "",
			args: args{
				head: []int{0, 1, 2},
				k:    4,
			},
			want: []int{2, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rotateRight(BuildList(tt.args.head), tt.args.k)
			assert.Equalf(t, tt.want, List2Slice(got), "rotateRight(%v, %v)", tt.args.head, tt.args.k)
		})
	}
}
