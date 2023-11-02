package _0xx

import (
	"testing"

	"github.com/stretchr/testify/require"
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
			want: []int{1, 2, 3, Null, 4},
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
			want: []int{1, 2, 3, Null, 4, Null, 5, Null, Null, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinTree2slice(tt.in)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_createTreeNodeFromSlice(t *testing.T) {
	tests := []struct {
		name string
		in   []int
	}{
		{
			name: "",
			in:   []int{1, 2, 3},
		},
		{
			name: "",
			in:   []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "",
			in:   []int{1, 2, 3, Null, 4, Null, 5, Null, Null, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateTreeNodeFromSlice(tt.in)
			gotSlice := BinTree2slice(got)
			require.Equal(t, tt.in, gotSlice)
		})
	}
}
