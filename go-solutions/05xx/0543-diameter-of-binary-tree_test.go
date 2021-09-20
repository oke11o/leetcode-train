package _5xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		tree []int
		want int
	}{
		{
			name: "",
			tree: []int{1, 2, 3, 4, 5},
			want: 3,
		},
		{
			name: "",
			tree: []int{1, 2},
			want: 1,
		},
		{
			name: "",
			tree: []int{1, 4, 3, 2, 5, nilTreeNodeVal, 8, nilTreeNodeVal, nilTreeNodeVal, 6, nilTreeNodeVal, 7, nilTreeNodeVal, nilTreeNodeVal, nilTreeNodeVal, 0, 9},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diameterOfBinaryTree(createTreeNodeFromSlice(tt.tree))
			require.Equal(t, tt.want, got)
		})
	}
}
