package _2xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_invertTree(t *testing.T) {
	tests := []struct {
		name string
		node []int
		want []int
	}{
		{
			name: "",
			node: []int{4, 2, 7, 1, 3, 6, 9},
			want: []int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			name: "",
			node: []int{2, 1, 3},
			want: []int{2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := invertTree(createTreeNodeFromSlice(tt.node))
			require.Equal(t, createTreeNodeFromSlice(tt.want), got)
		})
	}
}
