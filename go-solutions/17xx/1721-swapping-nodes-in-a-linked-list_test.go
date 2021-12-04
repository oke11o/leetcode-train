package _7xx

import (
	. "github.com/oke11o/leetcode-train/go-solutions/common"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_swapNodes(t *testing.T) {
	tests := []struct {
		name string
		head []int
		k    int
		want []int
	}{
		{
			name: "",
			head: []int{1, 2, 3, 4, 5},
			k:    2,
			want: []int{1, 4, 3, 2, 5},
		},
		{
			name: "",
			head: []int{7, 9, 6, 6, 7, 8, 3, 0, 9, 5},
			k:    5,
			want: []int{7, 9, 6, 6, 8, 7, 3, 0, 9, 5},
		},
		{
			name: "",
			head: []int{7, 9, 6, 6, 7, 8, 3, 0, 9, 5},
			k:    6,
			want: []int{7, 9, 6, 6, 8, 7, 3, 0, 9, 5},
		},
		{
			name: "",
			head: []int{1},
			k:    1,
			want: []int{1},
		},
		{
			name: "",
			head: []int{1, 2},
			k:    1,
			want: []int{2, 1},
		},
		{
			name: "",
			head: []int{1, 2, 3},
			k:    2,
			want: []int{1, 2, 3},
		},
		{
			name: "",
			head: []int{1, 2},
			k:    2,
			want: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := BuildList(tt.head)
			got := swapNodes(head, tt.k)
			require.Equal(t, tt.want, List2Slice(got))
		})
	}
}
