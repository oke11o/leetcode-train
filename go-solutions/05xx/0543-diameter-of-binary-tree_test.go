package _5xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/**
543. Diameter of Binary Tree
https://leetcode.com/problems/diameter-of-binary-tree/
Easy
#facebook #dfs
*/
func diameterOfBinaryTree_hardSolution(root *TreeNode) int {
	if root == nil {
		return 0
	}
	m := make(map[TreeNode]int)

	stack := []*TreeNode{root}

	diameter := 0

	for len(stack) > 0 {

		node := stack[len(stack)-1]

		leftOK := false
		if node.Left != nil {
			_, leftOK = m[*node.Left]
		}
		rightOK := false
		if node.Right != nil {
			_, rightOK = m[*node.Right]
		}
		if node.Left != nil && !leftOK {
			stack = append(stack, node.Left)
		} else if node.Right != nil && !rightOK {
			stack = append(stack, node.Right)
		} else {
			stack = stack[:len(stack)-1]
			leftDepth := 0
			if node.Left != nil {
				leftDepth = m[*node.Left]
			}
			rightDepth := 0
			if node.Right != nil {
				rightDepth = m[*node.Right]
			}

			max := leftDepth
			if max < rightDepth {
				max = rightDepth
			}
			max++
			m[*node] = max

			if diameter < leftDepth+rightDepth {
				diameter = leftDepth + rightDepth
			}

		}
	}

	return diameter
}

/*********************************/
/********** Solution 2 ***********/
/*********************************/
var diameterOfBinaryTreeResult int

func longestPath(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftPath := longestPath(node.Left)
	rightPath := longestPath(node.Right)

	if diameterOfBinaryTreeResult < leftPath+rightPath {
		diameterOfBinaryTreeResult = leftPath + rightPath
	}

	if leftPath > rightPath {
		return leftPath + 1
	}
	return rightPath + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	diameterOfBinaryTreeResult = 0
	longestPath(root)
	return diameterOfBinaryTreeResult
}

/*********************************/
/************* TESTS *************/
/*********************************/
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
			tree: []int{1, 4, 3, 2, 5, null, 8, null, null, 6, null, 7, null, null, null, 0, 9},
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
