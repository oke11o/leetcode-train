package _1xx

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
