package _1xx

import (
	. "github.com/oke11o/leetcode-train/go-solutions"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	tests := []struct {
		name string
		root []int
		want [][]int
	}{
		{
			name: "",
			root: []int{3, 9, 20, Null, Null, 15, 7},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name: "",
			root: []int{1},
			want: [][]int{{1}},
		},
		{
			name: "",
			root: []int{},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := levelOrder(CreateTreeNodeFromSlice(tt.root))
			require.Equal(t, tt.want, got)
		})
	}
}
