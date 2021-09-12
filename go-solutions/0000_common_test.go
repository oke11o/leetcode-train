package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_tree2slice(t *testing.T) {
	tests := []struct {
		name string
		in   *TreeNode
		want []int
	}{
		{
			name: "",
			in:   nil,
			want: nil,
		},
		{
			name: "",
			in: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			want: []int{1},
		},
		{
			name: "",
			in: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "",
			in: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: nil,
					Right: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
			want: []int{1, 2, 3, nilTreeNodeVal, 4},
		},
		{
			name: "",
			in: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: nil,
					Right: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:  3,
					Left: nil,
					Right: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val:   6,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: []int{1, 2, 3, nilTreeNodeVal, 4, nilTreeNodeVal, 5, nilTreeNodeVal, nilTreeNodeVal, nilTreeNodeVal, nilTreeNodeVal, nilTreeNodeVal, nilTreeNodeVal, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := binTree2slice(tt.in)
			require.Equal(t, tt.want, got)
		})
	}
}
func Test_createTreeNodeFromSlice(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want bool
	}{
		{
			name: "",
			in:   nil,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := createTreeNodeFromSlice(tt.in)
			require.Equal(t, tt.want, got)
		})
	}
}
