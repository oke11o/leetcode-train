package _0xx

var inorderTraversalResult []int

func travel(node *TreeNode) {
	if node == nil {
		return
	}
	travel(node.Left)
	inorderTraversalResult = append(inorderTraversalResult, node.Val)
	travel(node.Right)
}

// 0094. Binary Tree Inorder Traversal
// Easy
// Stack, Tree, DFS, Binary Tree
// https://leetcode.com/problems/binary-tree-inorder-traversal/
func inorderTraversal(root *TreeNode) []int {
	inorderTraversalResult = []int{}
	if root == nil {
		return inorderTraversalResult
	}
	travel(root)

	return inorderTraversalResult
}
