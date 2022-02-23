package _0xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_subsetsWithDup(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "",
			nums: []int{1, 2, 2},
			want: [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}},
		},
		{
			name: "",
			nums: []int{0},
			want: [][]int{{}, {0}},
		},
		{
			name: "",
			nums: []int{4, 4, 4, 1, 4},
			want: [][]int{{}, {1}, {1, 4}, {1, 4, 4}, {1, 4, 4, 4}, {1, 4, 4, 4, 4}, {4}, {4, 4}, {4, 4, 4}, {4, 4, 4, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsetsWithDup(tt.nums)
			require.ElementsMatch(t, tt.want, got)
		})
	}
}
