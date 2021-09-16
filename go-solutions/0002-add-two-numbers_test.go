package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

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
