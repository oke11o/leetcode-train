package _0xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

/*
2265. Count Nodes Equal to Average of Subtree
https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/
*/
func averageOfSubtree(root *TreeNode) int {
	result := 0
	// DFS
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		//if leaf
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil {
			result++
			return 1
		}
		leftChildren := dfs(node.Left)
		rightChildren := dfs(node.Right)

		cnt := leftChildren + rightChildren
		sum := 0
		if node.Left != nil {
			sum += node.Left.Val
		}
		if node.Right != nil {
			sum += node.Right.Val
		}
		sum += node.Val
		cnt += 1
		if (sum)/(cnt) == node.Val {
			result++
		}
		node.Val = sum
		return cnt
	}
	dfs(root)

	return result
}

func Test_averageOfSubtree(t *testing.T) {
	tests := []struct {
		name string
		root []int
		want int
	}{
		{
			name: "",
			root: []int{4, 8, 5, 0, 1, Null, 6},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := averageOfSubtree(CreateTreeNodeFromSlice(tt.root))
			require.Equal(t, got, tt.want)
		})
	}
}
