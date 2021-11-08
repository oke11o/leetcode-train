package _8xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_allPossibleFBT(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]int
	}{
		{
			name: "",
			n:    7,
			want: [][]int{{0, 0, 0, null, null, 0, 0, null, null, 0, 0}, {0, 0, 0, null, null, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, null, null, null, null, 0, 0}, {0, 0, 0, 0, 0, null, null, 0, 0}},
		},
		{
			name: "",
			n:    3,
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "",
			n:    20,
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := allPossibleFBT_withCache(tt.n)
			res := make([][]int, 0, len(got))
			for _, g := range got {
				res = append(res, binTree2slice(g))
			}
			require.Equal(t, tt.want, res)
		})
	}
}
