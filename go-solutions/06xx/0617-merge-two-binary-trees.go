package _6xx

func mergeTrees(node1 *TreeNode, node2 *TreeNode) *TreeNode {
	if node2 == nil {
		return node1
	}
	if node1 == nil {
		return node2
	}
	node1.Val += node2.Val
	node1.Left = mergeTrees(node1.Left, node2.Left)
	node1.Right = mergeTrees(node1.Right, node2.Right)
	return node1
}
