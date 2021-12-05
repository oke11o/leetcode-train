package _6xx

import "math"

const nilTreeNodeVal = math.MinInt32

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// nolint
func binTree2sliceRec(idx int, node *TreeNode, result []int) []int {
	for len(result) <= idx {
		result = append(result, nilTreeNodeVal)
	}
	val := nilTreeNodeVal
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
			result = append(result, nilTreeNodeVal)
		} else {
			result = append(result, node.Val)
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}
	lastNull := len(result) - 1
	for ; lastNull >= 0; lastNull-- {
		if result[lastNull] != nilTreeNodeVal {
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
		if val != nilTreeNodeVal {
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
