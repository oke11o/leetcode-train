package _1xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_hasPathSum(t *testing.T) {
	tests := []struct {
		name      string
		root      []int
		targetSum int
		want      bool
	}{
		{
			name:      "",
			root:      []int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, null, 1},
			targetSum: 22,
			want:      true,
		},
		{
			name:      "",
			root:      []int{1, 2, 3},
			targetSum: 2,
			want:      false,
		},
		{
			name:      "",
			root:      []int{1, 2},
			targetSum: 0,
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasPathSum(createTreeNodeFromSlice(tt.root), tt.targetSum)
			require.Equal(t, tt.want, got)
		})
	}
}
