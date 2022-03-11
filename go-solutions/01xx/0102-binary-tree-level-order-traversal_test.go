package _1xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// 0102. Binary Tree Level Order Traversal
// Medium
// Tree, BFS, Binary Tree
// https://leetcode.com/problems/binary-tree-level-order-traversal/
// Amazon - https://www.educative.io/blog/crack-amazon-coding-interview-questions
func levelOrder_hardSolution(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	levelIdx := 0
	pushIdx := 1
	queues := [][]*TreeNode{{root}}
	result := [][]int{}

	for len(queues[levelIdx]) > 0 {
		if pushIdx > len(queues)-1 {
			queues = append(queues, nil)
		}
		if levelIdx > len(result)-1 {
			result = append(result, nil)
		}
		cur := queues[levelIdx][0]
		queues[levelIdx] = queues[levelIdx][1:]
		result[levelIdx] = append(result[levelIdx], cur.Val)

		if cur.Left != nil {
			queues[pushIdx] = append(queues[pushIdx], cur.Left)
		}
		if cur.Right != nil {
			queues[pushIdx] = append(queues[pushIdx], cur.Right)
		}
		if len(queues[levelIdx]) == 0 {
			levelIdx++
			pushIdx++
		}
	}

	return result
}

/**
recursive
*/
func levelOrder(root *TreeNode) [][]int {
	levels := [][]int{}
	if root == nil {
		return levels
	}
	var travel func(node *TreeNode, level int)
	travel = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level > len(levels)-1 {
			levels = append(levels, []int{})
		}
		levels[level] = append(levels[level], node.Val)
		travel(node.Left, level+1)
		travel(node.Right, level+1)
	}

	travel(root, 0)

	return levels
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_levelOrder(t *testing.T) {
	tests := []struct {
		name string
		root []int
		want [][]int
	}{
		{
			name: "",
			root: []int{3, 9, 20, Null, Null, 15, 7},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name: "",
			root: []int{1},
			want: [][]int{{1}},
		},
		{
			name: "",
			root: []int{},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := levelOrder(CreateTreeNodeFromSlice(tt.root))
			require.Equal(t, tt.want, got)
		})
	}
}
