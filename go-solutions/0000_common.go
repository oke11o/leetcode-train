package go_solutions

const nilTreeNodeVal = -1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2slice(in *TreeNode) []int {
	var slFn = func() {

	}
	_ = slFn
	if in == nil {
		return nil
	}
	result := []int{1}

	return result
}

func createTreeNodeFromSlice(in []int) *TreeNode {
	return nil
}

func sprintTreeNode(node *TreeNode) string {
	return ""
}
