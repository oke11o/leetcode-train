package _2xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeElements(t *testing.T) {
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
		head []int
		val  int
		want []int
	}{
		{
			name: "",
			head: []int{1, 2, 6, 3, 4, 5, 6},
			val:  6,
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "",
			head: []int{},
			val:  1,
			want: []int{},
		},
		{
			name: "",
			head: []int{7, 7, 7, 7},
			val:  7,
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElements(buildList(tt.head), tt.val)
			want := buildList(tt.want)
			require.Equal(t, want, got)
		})
	}
}
