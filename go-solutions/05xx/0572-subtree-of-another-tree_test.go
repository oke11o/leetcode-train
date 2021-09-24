package _5xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_isSubtree(t *testing.T) {
	tests := []struct {
		name    string
		root    []int
		subRoot []int
		want    bool
	}{
		{
			name:    "",
			root:    []int{3, 4, 5, 1, 2},
			subRoot: []int{4, 1, 2},
			want:    true,
		},
		{
			name:    "",
			root:    []int{3, 4, 5, 1, 2, null, null, null, null, 0},
			subRoot: []int{4, 1, 2},
			want:    false,
		},
		{
			name:    "",
			root:    []int{1, null, 1, null, 1, null, 1, null, 1, null, 1, null, 1, null, 1, 2},
			subRoot: []int{1, null, 1, null, 1, null, 1, null, 1, null, 1, 2},
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSubtree(createTreeNodeFromSlice(tt.root), createTreeNodeFromSlice(tt.subRoot))
			require.Equal(t, tt.want, got)
		})
	}
}
