package go_solutions

func isValidBST(node *TreeNode) bool {
	if node == nil {
		return false
	}
	if node.Left != nil && node.Right != nil {
		if !(node.Left.Val < node.Val && node.Val < node.Right.Val) {
			return false
		}
		return isValidBST(node.Left) && isValidBST(node.Right)
	}
	if node.Left == nil && node.Right != nil {
		if !(node.Val < node.Right.Val) {

		}
		return isValidBST(node.Right)
	}
	return isValidBST(node.Left)
}
