package _5xx

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	m := make(map[TreeNode]int)

	stack := []*TreeNode{root}

	diameter := 0

	for len(stack) > 0 {

		node := stack[len(stack)-1]

		leftOK := false
		if node.Left != nil {
			_, leftOK = m[*node.Left]
		}
		rightOK := false
		if node.Right != nil {
			_, rightOK = m[*node.Left]
		}
		if node.Left != nil && !leftOK {
			stack = append(stack, node.Left)
		} else if node.Right != nil && !rightOK {
			stack = append(stack, node.Right)
		} else {
			stack = stack[:len(stack)-1]
			leftDepth := 0
			if node.Left != nil {
				leftDepth = m[*node.Left]
			}
			rightDepth := 0
			if node.Left != nil {
				rightDepth = m[*node.Right]
			}

			max := leftDepth
			if max < rightDepth {
				max = rightDepth
			}
			max++
			m[*node] = max

			if diameter < leftDepth+rightDepth {
				diameter = leftDepth + rightDepth
			}

		}
	}

	return diameter + 1
}
