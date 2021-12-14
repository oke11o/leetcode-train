package _0xx

import (
	. "github.com/oke11o/leetcode-train/go-solutions"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	tests := []struct {
		name string
		root []int
		want bool
	}{
		{
			name: "",
			root: []int{2, 1, 3},
			want: true,
		},
		{
			name: "",
			root: []int{5, 1, 4, Null, Null, 3, 6},
			want: false,
		},
		{
			name: "",
			root: []int{2, 2, 2},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidBST(CreateTreeNodeFromSlice(tt.root))
			require.Equal(t, tt.want, got)
		})
	}
}
