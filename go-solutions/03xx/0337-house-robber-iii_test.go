package _3xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_rob(t *testing.T) {
	tests := []struct {
		name string
		root []int
		want int
	}{
		{
			name: "",
			root: []int{3, 2, 3, null, 3, null, 1},
			want: 7,
		},
		{
			name: "",
			root: []int{3, 4, 5, 1, 3, null, 1},
			want: 9,
		},
		{
			name: "",
			root: []int{3, 2, 3, 4, 3, null, 1, 5},
			want: 12,
		},
		{
			name: "",
			root: []int{4, 1, null, 2, null, 3},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rob(createTreeNodeFromSlice(tt.root))
			require.Equal(t, tt.want, got)
		})
	}
}
