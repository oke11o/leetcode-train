package _1xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	tests := []struct {
		name string
		tree []int
		want int
	}{
		{
			name: "",
			tree: []int{3, 9, 20, nilTreeNodeVal, nilTreeNodeVal, 15, 7},
			want: 3,
		},
		{
			name: "",
			tree: []int{1, nilTreeNodeVal, 2},
			want: 2,
		},
		{
			name: "",
			tree: []int{},
			want: 0,
		},
		{
			name: "",
			tree: []int{0},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxDepth(createTreeNodeFromSlice(tt.tree))
			require.Equal(t, tt.want, got)
		})
	}
}
