package _1xx

// 0112. Path Sum
func hasPathSum(node *TreeNode, targetSum int) bool {
	if node == nil {
		return false
	}

	if node.Left == nil && node.Right == nil {
		if node.Val == targetSum {
			return true
		} else {
			return false
		}
	}

	return hasPathSum(node.Left, targetSum-node.Val) || hasPathSum(node.Right, targetSum-node.Val)
}
