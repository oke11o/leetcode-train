package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_hasCycle(t *testing.T) {
	var buildList = func(list []*ListNode, last int) *ListNode {
		if len(list) == 0 {
			return nil
		}
		for i := 1; i < len(list); i++ {
			list[i-1].Next = list[i]
		}
		if last >= 0 {
			list[len(list)-1].Next = list[last]
		}
		return list[0]
	}

	tests := []struct {
		name string
		list []*ListNode
		last int
		want bool
	}{
		{
			name: "",
			list: []*ListNode{{Val: 3}, {Val: 2}, {Val: 0}, {Val: -4}},
			last: 1,
			want: true,
		},
		{
			name: "",
			list: []*ListNode{{Val: 1}, {Val: 2}},
			last: 0,
			want: true,
		},
		{
			name: "",
			list: []*ListNode{{Val: 1}},
			last: -1,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := hasCycle(buildList(tt.list, tt.last))
			require.Equal(t, tt.want, got)
		})
	}
}