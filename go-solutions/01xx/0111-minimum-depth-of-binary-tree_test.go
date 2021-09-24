package _1xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_minDepth(t *testing.T) {
	tests := []struct {
		name string
		root []int
		want int
	}{
		{
			name: "",
			root: []int{3, 9, 20, null, null, 15, 7},
			want: 2,
		},
		{
			name: "",
			root: []int{2, null, 3, null, 4, null, 5},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minDepth(createTreeNodeFromSlice(tt.root))
			require.Equal(t, tt.want, got)
		})
	}
}
