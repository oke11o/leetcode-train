package _1xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxDepth(t *testing.T) {
	tests := []struct {
		name string
		tree []int
		want int
	}{
		{
			name: "",
			tree: []int{3, 9, 20, Null, Null, 15, 7},
			want: 3,
		},
		{
			name: "",
			tree: []int{1, Null, 2},
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
			got := maxDepth(CreateTreeNodeFromSlice(tt.tree))
			require.Equal(t, tt.want, got)
		})
	}
}
