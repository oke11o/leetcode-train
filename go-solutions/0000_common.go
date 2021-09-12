package go_solutions

import "math"

const nilTreeNodeVal = math.MinInt

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func binTree2slice(in *TreeNode) []int {
	if in == nil {
		return nil
	}

	return binTree2sliceRec(0, in, nil)
}

func createTreeNodeFromSlice(in []int) *TreeNode {
	return nil
}

func sprintTreeNode(node *TreeNode) string {
	return ""
}
