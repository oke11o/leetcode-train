package _1xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_delNodes(t *testing.T) {
	tests := []struct {
		name      string
		root      []int
		to_delete []int
		want      [][]int
	}{
		{
			name:      "",
			root:      []int{1, 2, 3, 4, 5, 6, 7},
			to_delete: []int{3, 5},
			want:      [][]int{{1, 2, null, 4}, {6}, {7}},
		},
		{
			name:      "",
			root:      []int{1, 2, 4, null, 3},
			to_delete: []int{3},
			want:      [][]int{{1, 2, 4}},
		},
		{
			name:      "Root doesn't pass. Child pass full",
			root:      []int{1, 2, 3, null, null, null, 4},
			to_delete: []int{2, 1},
			want:      [][]int{{3, null, 4}},
		},
		{
			name:      "У ребенка на удаление есть еще дети на удаление",
			root:      []int{1, 2, null, 4, 3},
			to_delete: []int{2, 3},
			want:      [][]int{{1}, {4}},
		},
		{
			name:      "У ребенка на удаление есть еще дети на удаление (HARD)",
			root:      []int{1, 2, null, 4, 3, 5, 6, 7, 8, null, null, null, null, 9, 10},
			to_delete: []int{2, 7},
			want:      [][]int{{1}, {4, 5, 6}, {3, null, 8}, {9}, {10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := delNodes(createTreeNodeFromSlice(tt.root), tt.to_delete)
			var want []*TreeNode
			for _, w := range tt.want {
				want = append(want, createTreeNodeFromSlice(w))
			}
			require.Equal(t, want, got)
		})
	}
}
