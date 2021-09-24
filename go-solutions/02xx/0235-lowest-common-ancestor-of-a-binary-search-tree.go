package _2xx

// 0235. Lowest Common Ancestor of a Binary Search Tree
func lowestCommonAncestor(node, p, q *TreeNode) *TreeNode {
	for node != nil {
		if p.Val < node.Val && q.Val < node.Val {
			node = node.Left
		} else if p.Val > node.Val && q.Val > node.Val {
			node = node.Right
		} else {
			return node
		}
	}
	return node
}
