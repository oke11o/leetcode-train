package _1xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_isSameTree(t *testing.T) {
	tests := []struct {
		name string
		p    []int
		q    []int
		want bool
	}{
		{
			name: "",
			p:    []int{1, 2, 3},
			q:    []int{1, 2, 3},
			want: true,
		},
		{
			name: "",
			p:    []int{1, 2},
			q:    []int{1, null, 2},
			want: false,
		},
		{
			name: "",
			p:    []int{1, 2, 1},
			q:    []int{1, 1, 2},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := createTreeNodeFromSlice(tt.p)
			q := createTreeNodeFromSlice(tt.q)

			got := isSameTree(p, q)
			require.Equal(t, tt.want, got)
		})
	}
}
