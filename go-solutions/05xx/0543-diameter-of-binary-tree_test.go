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
			tree: []int{1, 2, 3, 4, 5, nilTreeNodeVal, 7, 8, 9, 10, 11, nilTreeNodeVal, nilTreeNodeVal, nilTreeNodeVal, nilTreeNodeVal, 16, 17, 18, 19, 20, 21, 22, 23},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diameterOfBinaryTree(createTreeNodeFromSlice(tt.tree, 0))
			require.Equal(t, tt.want, got)
		})
	}
}
