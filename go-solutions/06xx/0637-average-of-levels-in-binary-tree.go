package _6xx

// 0637. Average of Levels in Binary Tree
func averageOfLevels(root *TreeNode) []float64 {
	queue := []*TreeNode{root}

	var result []float64
	for len(queue) > 0 {

		tmpQ := []*TreeNode{}
		levelSum := 0
		for _, n := range queue {
			levelSum += (n.Val)
			if n.Left != nil {
				tmpQ = append(tmpQ, n.Left)
			}
			if n.Right != nil {
				tmpQ = append(tmpQ, n.Right)
			}
		}
		result = append(result, float64(levelSum)/float64(len(queue)))
		queue = tmpQ
	}

	return result
}
