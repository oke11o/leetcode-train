package _1xx

import (
	. "github.com/oke11o/leetcode-train/go-solutions"
)

// 0102. Binary Tree Level Order Traversal
// Medium
// Tree, BFS, Binary Tree
// https://leetcode.com/problems/binary-tree-level-order-traversal/
// Amazon - https://www.educative.io/blog/crack-amazon-coding-interview-questions
func levelOrder(root *TreeNode) [][]int {
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
