package _1xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "",
			nums: []int{-10, -3, 0, 5, 9},
			want: []int{0, -3, 9, -10, null, 5},
		},
		{
			name: "",
			nums: []int{1, 3},
			want: []int{3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedArrayToBST(tt.nums)
			require.Equal(t, createTreeNodeFromSlice(tt.want), got)
		})
	}
}
