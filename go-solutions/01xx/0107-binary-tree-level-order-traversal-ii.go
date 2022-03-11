package _1xx

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	qIdx := 0
	nextIdx := 1
	queues := [][]*TreeNode{{root}}
	result := [][]int{}

	for len(queues[qIdx]) > 0 {
		if nextIdx > len(queues)-1 {
			queues = append(queues, nil)
		}
		if qIdx > len(result)-1 {
			result = append(result, nil)
		}
		cur := queues[qIdx][0]
		queues[qIdx] = queues[qIdx][1:]
		result[qIdx] = append(result[qIdx], cur.Val)

		if cur.Left != nil {
			queues[nextIdx] = append(queues[nextIdx], cur.Left)
		}
		if cur.Right != nil {
			queues[nextIdx] = append(queues[nextIdx], cur.Right)
		}
		if len(queues[qIdx]) == 0 {
			qIdx++
			nextIdx++
		}
	}
	r := [][]int{}
	for i := len(result) - 1; i >= 0; i-- {
		r = append(r, result[i])
	}

	return r

}
