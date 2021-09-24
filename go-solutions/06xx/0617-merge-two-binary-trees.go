package _6xx

func mergeTrees(node1 *TreeNode, node2 *TreeNode) *TreeNode {
	if node1 == nil && node2 == nil {
		return nil
	}
	if node1 != nil && node2 == nil {
		node1.Left = mergeTrees(node1.Left, nil)
		node1.Right = mergeTrees(node1.Right, nil)
		return node1
	}
	if node2 != nil && node1 == nil {
		node2.Left = mergeTrees(node2.Left, nil)
		node2.Right = mergeTrees(node2.Right, nil)
		return node2
	}
	node1.Val = node1.Val + node2.Val
	node1.Left = mergeTrees(node1.Left, node2.Left)
	node1.Right = mergeTrees(node1.Right, node2.Right)

	return node1
}
