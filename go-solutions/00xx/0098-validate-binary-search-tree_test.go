package _0xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/**
https://leetcode.com/problems/validate-binary-search-tree/
0098. Validate Binary Search Tree
Medium
DFS
*/
func isValidBST(node *TreeNode) bool {
	var traverseAndCheck func(node *TreeNode, min, max int) bool
	traverseAndCheck = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		if node.Val <= min || node.Val >= max {
			return false
		}
		return traverseAndCheck(node.Left, min, node.Val) && traverseAndCheck(node.Right, node.Val, max)
	}

	maxInt := 1<<63 - 1
	minInt := -1 << 63
	return traverseAndCheck(node, minInt, maxInt)
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_isValidBST(t *testing.T) {
	tests := []struct {
		name string
		root []int
		want bool
	}{
		{
			name: "",
			root: []int{2, 1, 3},
			want: true,
		},
		{
			name: "",
			root: []int{5, 1, 4, Null, Null, 3, 6},
			want: false,
		},
		{
			name: "",
			root: []int{2, 2, 2},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidBST(CreateTreeNodeFromSlice(tt.root))
			require.Equal(t, tt.want, got)
		})
	}
}
