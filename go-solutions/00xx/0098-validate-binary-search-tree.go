package _0xx

import (
	. "github.com/oke11o/leetcode-train/go-solutions"
)

func isValidBSTTraverseAndCheck(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	return isValidBSTTraverseAndCheck(node.Left, min, node.Val) && isValidBSTTraverseAndCheck(node.Right, node.Val, max)
}

// 0098. Validate Binary Search Tree
// Medium
// DFS
// https://leetcode.com/problems/validate-binary-search-tree/
func isValidBST(node *TreeNode) bool {
	intSize := 32 << (^uint(0) >> 63) // 32 or 64
	MaxInt := 1<<(intSize-1) - 1
	MinInt := -1 << (intSize - 1)

	return isValidBSTTraverseAndCheck(node, MinInt, MaxInt)
}
