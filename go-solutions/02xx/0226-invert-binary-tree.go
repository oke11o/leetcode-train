package _2xx

// 0226. Invert Binary Tree
func invertTree(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	node.Left, node.Right = invertTree(node.Right), invertTree(node.Left)
	return node
}
