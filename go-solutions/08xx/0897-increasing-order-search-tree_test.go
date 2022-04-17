package _8xx

import (
	"reflect"
	"testing"
)

/**
https://leetcode.com/problems/increasing-order-search-tree/
897. Increasing Order Search Tree
Easy
*/

func increasingBST(root *TreeNode) *TreeNode {
	dummy := &TreeNode{}
	var increase func(node *TreeNode)
	increase = func(node *TreeNode) {
		if node == nil {
			return
		}
		increase(node.Left)
		node.Left = nil
		dummy.Right = node
		dummy = dummy.Right
		increase(node.Right)
	}
	result := dummy
	increase(root)
	return result.Right
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_increasingBST(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				root: []int{5, 3, 6, 2, 4, null, 8, 1, null, null, null, 7, 9},
			},
			want: []int{1, null, 2, null, 3, null, 4, null, 5, null, 6, null, 7, null, 8, null, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := increasingBST(createTreeNodeFromSlice(tt.args.root))
			if !reflect.DeepEqual(binTree2slice(got), tt.want) {
				t.Errorf("increasingBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
