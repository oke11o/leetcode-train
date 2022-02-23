package _2xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPalindrome(t *testing.T) {
	var buildList = func(in []int) *ListNode {
		if len(in) == 0 {
			return nil
		}
		head := &ListNode{
			Val: in[0],
		}
		result := head
		for i := 1; i < len(in); i++ {
			item := &ListNode{Val: in[i]}
			head.Next = item
			head = item
		}

		return result
	}

	tests := []struct {
		name string
		list []int
		want bool
	}{
		{
			name: "",
			list: []int{1, 2, 2, 1},
			want: true,
		},
		{
			name: "",
			list: []int{1, 2},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(buildList(tt.list))
			require.Equal(t, tt.want, got)
		})
	}
}
