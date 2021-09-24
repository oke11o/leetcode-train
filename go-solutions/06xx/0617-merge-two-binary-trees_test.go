package _6xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_mergeTrees(t *testing.T) {
	tests := []struct {
		name  string
		node1 []int
		node2 []int
		want  []int
	}{
		{
			name:  "",
			node1: []int{1, 3, 2, 5},
			node2: []int{2, 1, 3, nilTreeNodeVal, 4, nilTreeNodeVal, 7},
			want:  []int{3, 4, 5, 5, 4, nilTreeNodeVal, 7},
		},
		{
			name:  "",
			node1: []int{1},
			node2: []int{1, 2},
			want:  []int{2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTrees(createTreeNodeFromSlice(tt.node1), createTreeNodeFromSlice(tt.node2))
			require.Equal(t, createTreeNodeFromSlice(tt.want), got)
		})
	}
}
