package _1xx

// 1110. Delete Nodes And Return Forest
func delNodes(root *TreeNode, toDelete []int) []*TreeNode {
	var result []*TreeNode
	mapa := make(map[int]struct{}, len(toDelete))
	for _, v := range toDelete {
		mapa[v] = struct{}{}
	}
	queue := []*TreeNode{root}
	if _, ok := mapa[root.Val]; !ok {
		result = append(result, root)
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if _, ok := mapa[node.Val]; ok {
			if node.Left != nil {
				if _, ok := mapa[node.Left.Val]; !ok {
					result = append(result, node.Left)
				}
			}
			if node.Right != nil {
				if _, ok := mapa[node.Right.Val]; !ok {
					result = append(result, node.Right)
				}
			}
		}

		// Детей надо в любом случае добавить в обход
		if node.Left != nil {
			queue = append(queue, node.Left)
			if _, ok := mapa[node.Left.Val]; ok {
				node.Left = nil
			}
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			if _, ok := mapa[node.Right.Val]; ok {
				node.Right = nil
			}
		}
	}

	return result
}

// TODO: попробовать с dfs. Например: https://zhenyu0519.github.io/2020/03/21/lc1110/
