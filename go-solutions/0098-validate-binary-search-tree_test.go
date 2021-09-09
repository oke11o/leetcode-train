package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidBST(tt.root)
			require.Equal(t, tt.want, got)
		})
	}
}
