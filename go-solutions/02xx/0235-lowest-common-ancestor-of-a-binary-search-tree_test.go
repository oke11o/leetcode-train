package _2xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name string
		root []int
		p    []int
		q    []int
		want []int
	}{
		{
			name: "",
			root: []int{6, 2, 8, 0, 4, 7, 9, null, null, 3, 5},
			p:    []int{2, 0, 4, null, null, 3, 5},
			q:    []int{8, 7, 9},
			want: []int{6, 2, 8, 0, 4, 7, 9, null, null, 3, 5},
		},
		{
			name: "",
			root: []int{6, 2, 8, 0, 4, 7, 9, null, null, 3, 5},
			p:    []int{2, 0, 4, null, null, 3, 5},
			q:    []int{4, 3, 5},
			want: []int{2, 0, 4, null, null, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lowestCommonAncestor(createTreeNodeFromSlice(tt.root), createTreeNodeFromSlice(tt.p), createTreeNodeFromSlice(tt.q))
			require.Equal(t, createTreeNodeFromSlice(tt.want), got)
		})
	}
}
