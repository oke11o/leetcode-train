package _1xx

import (
	. "github.com/oke11o/leetcode-train/go-solutions"
)

func maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := maxDepth(node.Left)
	right := maxDepth(node.Right)
	if right < left {
		right = left
	}
	return right + 1
}
