package _1xx

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSymmetric(root *TreeNode) bool {
	queue := []*TreeNode{root, root}
	for len(queue) != 0 {
		n1 := queue[0]
		n2 := queue[1]
		queue = queue[2:]
		if n1 == nil && n2 == nil {
			continue
		}
		if n1 == nil || n2 == nil {
			return false
		}
		if n1.Val != n2.Val {
			return false
		}
		queue = append(queue, n1.Left)
		queue = append(queue, n2.Right)
		queue = append(queue, n1.Right)
		queue = append(queue, n2.Left)
	}
	return true
}

func isSymmetric_recursive(root *TreeNode) bool {
	var isMirrow func(n1, n2 *TreeNode) bool
	isMirrow = func(n1, n2 *TreeNode) bool {
		if n1 == nil && n2 == nil {
			return true
		}
		if n1 == nil || n2 == nil {
			return false
		}
		return n1.Val == n2.Val &&
			isMirrow(n1.Left, n2.Right) &&
			isMirrow(n1.Right, n2.Left)
	}
	return isMirrow(root, root)
}
