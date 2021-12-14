package _1xx

import (
	. "github.com/oke11o/leetcode-train/go-solutions"
)

// 0111. Minimum Depth of Binary Tree
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	queue := []*TreeNode{root}
Loop:
	for len(queue) != 0 {
		level++
		var tmpQ []*TreeNode
		for _, t := range queue {
			if t.Left == nil && t.Right == nil {
				break Loop
			}
			if t.Left != nil {
				tmpQ = append(tmpQ, t.Left)
			}
			if t.Right != nil {
				tmpQ = append(tmpQ, t.Right)
			}
		}
		queue = tmpQ
	}
	return level
}
