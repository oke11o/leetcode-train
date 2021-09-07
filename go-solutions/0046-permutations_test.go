package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_permute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "",
			nums: []int{1},
			want: [][]int{{1}},
		},
		{
			name: "",
			nums: []int{2, 1},
			want: [][]int{{2, 1}, {1, 2}},
		},
		{
			name: "",
			nums: []int{3, 2, 1},
			want: [][]int{
				{3, 2, 1}, {3, 1, 2},
				{2, 3, 1}, {2, 1, 3},
				{1, 2, 3}, {1, 3, 2},
			},
		},
		{
			name: "",
			nums: []int{1, 2, 3, 4},
			want: [][]int{
				{1, 2, 3, 4}, {1, 2, 4, 3}, {1, 3, 2, 4}, {1, 3, 4, 2}, {1, 4, 2, 3}, {1, 4, 3, 2},
				{2, 1, 3, 4}, {2, 1, 4, 3}, {2, 3, 1, 4}, {2, 3, 4, 1}, {2, 4, 1, 3}, {2, 4, 3, 1},
				{3, 1, 2, 4}, {3, 1, 4, 2}, {3, 2, 1, 4}, {3, 2, 4, 1}, {3, 4, 1, 2}, {3, 4, 2, 1},
				{4, 1, 2, 3}, {4, 1, 3, 2}, {4, 2, 1, 3}, {4, 2, 3, 1}, {4, 3, 1, 2}, {4, 3, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
