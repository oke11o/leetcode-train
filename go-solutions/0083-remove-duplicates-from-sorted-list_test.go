package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
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
		list []int
		want []int
	}{
		{
			name: "",
			list: []int{1, 1, 2},
			want: []int{1, 2},
		},
		{
			name: "",
			list: []int{1, 1, 2, 3, 3},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteDuplicates(buildList(tt.list))
			want := buildList(tt.want)
			require.Equal(t, want, got)
		})
	}
}
