package _1xx

// 1110. Delete Nodes And Return Forest
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	var treeInSlice = func(node *TreeNode, list []int) bool {
		if node == nil {
			return false
		}
		for _, val := range list {
			if node.Val == val {
				return true
			}
		}
		return false
	}
	if treeInSlice(root, to_delete) {
		return []*TreeNode{nil, root}
	}

	queue := []*TreeNode{root}
	orphans := []*TreeNode{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			if treeInSlice(node.Left, to_delete) {
				if node.Left.Left != nil && !treeInSlice(node.Left.Left, to_delete) {
					orphans = append(orphans, node.Left.Left)
				}
				if node.Left.Right != nil && !treeInSlice(node.Left.Right, to_delete) {
					orphans = append(orphans, node.Left.Right)
				}
				node.Left = nil
			} else {
				queue = append(queue, node.Left)
			}
		}
		if node.Right != nil {
			if treeInSlice(node.Right, to_delete) {
				if node.Right.Left != nil && !treeInSlice(node.Right.Left, to_delete) {
					orphans = append(orphans, node.Right.Left)
				}
				if node.Right.Right != nil && !treeInSlice(node.Right.Right, to_delete) {
					orphans = append(orphans, node.Right.Right)
				}
				node.Right = nil
			} else {
				queue = append(queue, node.Right)
			}
		}
	}

	result := []*TreeNode{root}
	if len(orphans) > 0 {
		result = append(result, orphans...)
	}
	return result
}
