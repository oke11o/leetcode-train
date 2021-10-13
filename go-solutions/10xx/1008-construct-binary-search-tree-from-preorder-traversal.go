package _0xx

// 1008. Construct Binary Search Tree from Preorder Traversal
func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	node := &TreeNode{Val: preorder[0]}
	preorder = preorder[1:]
	n := len(preorder)
	for i, t := range preorder {
		if t > node.Val {
			n = i
			break
		}
	}
	node.Left = bstFromPreorder(preorder[0:n])
	node.Right = bstFromPreorder(preorder[n:])

	return node
}
