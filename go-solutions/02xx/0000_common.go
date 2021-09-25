package _2xx

import "math"

const null = math.MinInt32

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binTree2sliceRec(idx int, node *TreeNode, result []int) []int {
	for len(result) <= idx {
		result = append(result, null)
	}
	val := null
	var left, right *TreeNode
	if node != nil {
		val = node.Val
		left = node.Left
		right = node.Right
	}
	result[idx] = val
	if left != nil {
		result = binTree2sliceRec(idx*2+1, left, result)
	}
	if right != nil {
		result = binTree2sliceRec(idx*2+2, right, result)
	}

	return result
}

func binTree2slice(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			result = append(result, null)
		} else {
			result = append(result, node.Val)
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}
	lastNull := len(result) - 1
	for ; lastNull >= 0; lastNull-- {
		if result[lastNull] != null {
			break
		}
	}

	return result[:lastNull+1]
}

func createTreeNodeFromSlice(in []int) *TreeNode {
	if len(in) == 0 {
		return nil
	}

	root := &TreeNode{Val: in[0]}

	nodeList := []*TreeNode{root}

	nodeIdx := 0
	for i := 1; i < len(in); i++ {
		val := in[i]
		var node *TreeNode
		if val != null {
			node = &TreeNode{Val: val}
		}
		nodeList = append(nodeList, node)
		if i%2 == 1 {
			nodeList[nodeIdx].Left = node
		} else {
			nodeList[nodeIdx].Right = node

			nodeIdx++
			for nodeIdx < len(nodeList)-1 && nodeList[nodeIdx] == nil {
				nodeIdx++
			}
		}
	}

	return root
}

func sprintTreeNode(node *TreeNode) string {
	return ""
}
