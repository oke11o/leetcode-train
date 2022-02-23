package _0xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_inorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		root []int
		want []int
	}{
		{
			name: "",
			root: []int{1, Null, 2, 3},
			want: []int{1, 3, 2},
		},
		{
			name: "",
			root: []int{1},
			want: []int{1},
		},
		{
			name: "",
			root: []int{},
			want: []int{},
		},
		{
			name: "",
			root: []int{1, Null, 2},
			want: []int{1, 2},
		},
		{
			name: "",
			root: []int{1, 2},
			want: []int{2, 1},
		},
		{
			name: "",
			root: []int{10, 5, 20, 3, 7, 35},
			want: []int{3, 5, 7, 10, 35, 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inorderTraversal(CreateTreeNodeFromSlice(tt.root))
			require.Equal(t, tt.want, got)
		})
	}
}
