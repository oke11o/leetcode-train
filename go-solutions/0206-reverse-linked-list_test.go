package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_reverseList(t *testing.T) {
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
		want []int
	}{
		{
			name: "",
			list: []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "",
			list: []int{1, 2},
			want: []int{2, 1},
		},
		{
			name: "",
			list: []int{},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			rev := reverseList(buildList(tt.list))
			if len(tt.want) == 0 {
				require.Nil(t, rev)
				return
			}

			var got []int
			for rev.Next != nil {
				got = append(got, rev.Val)
				rev = rev.Next
			}
			got = append(got, rev.Val)
			require.Equal(t, tt.want, got)
		})
	}
}
