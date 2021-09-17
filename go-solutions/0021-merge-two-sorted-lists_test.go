package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
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
			l1:   []int{1, 2, 4},
			l2:   []int{1, 3, 4},
			want: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name: "",
			l1:   []int{-9, 3},
			l2:   []int{5, 7},
			want: []int{-9, 3, 5, 7},
		},
		{
			name: "",
			l1:   []int{},
			l2:   []int{},
			want: []int{},
		},
		{
			name: "",
			l1:   []int{},
			l2:   []int{0},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(buildList(tt.l1), buildList(tt.l2))
			want := buildList(tt.want)
			require.Equal(t, want, got)
		})
	}
}
