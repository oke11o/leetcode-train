package _3xx

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 0337. House Robber III
// Состояние динамики [withRoot, withoutRoot]
func rob(root *TreeNode) int {
	rs := rob_state(root)
	return max(rs.withRoot, rs.withoutRoot)
}

type robState struct {
	withRoot, withoutRoot int
}

func rob_state(node *TreeNode) robState {
	if node == nil {
		return robState{}
	}
	leftState := rob_state(node.Left)
	rightState := rob_state(node.Right)
	withRoot := node.Val + leftState.withoutRoot + rightState.withoutRoot
	withoutRoot := max(leftState.withRoot, leftState.withoutRoot) + max(rightState.withRoot, rightState.withoutRoot) // тут надо брать максимум, так как у нас нет строгого указания, что можно брать только с корнем. Может у ребенка без корня все равно лучше результат. Не надо тут ребенка ограничивать

	return robState{withRoot, withoutRoot}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
