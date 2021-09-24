package _5xx

// 0572. Subtree of Another Tree
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	if isSameTree_Rec(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree_Rec(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return isSameTree_Rec(p.Left, q.Left) && isSameTree_Rec(p.Right, q.Right)
}
