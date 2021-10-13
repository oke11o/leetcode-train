package _0xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_bstFromPreorder(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		want     *TreeNode
	}{
		{
			name:     "",
			preorder: []int{8, 5, 1, 7, 10, 12},
			want:     createTreeNodeFromSlice([]int{8, 5, 10, 1, 7, null, 12}),
		},
		{
			name:     "",
			preorder: []int{4, 2},
			want:     createTreeNodeFromSlice([]int{4, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bstFromPreorder(tt.preorder)
			require.Equal(t, tt.want, got)
		})
	}
}
