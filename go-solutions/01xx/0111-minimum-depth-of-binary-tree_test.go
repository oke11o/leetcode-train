package _1xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minDepth(t *testing.T) {
	tests := []struct {
		name string
		root []int
		want int
	}{
		{
			name: "",
			root: []int{3, 9, 20, Null, Null, 15, 7},
			want: 2,
		},
		{
			name: "",
			root: []int{2, Null, 3, Null, 4, Null, 5},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minDepth(CreateTreeNodeFromSlice(tt.root))
			require.Equal(t, tt.want, got)
		})
	}
}
