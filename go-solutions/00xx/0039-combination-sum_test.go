package _0xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
		wan1       [][]int
	}{

		{
			name:       "",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			want:       [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "",
			candidates: []int{2, 3, 5},
			target:     8,
			want:       [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "",
			candidates: []int{2},
			target:     1,
			want:       [][]int{},
		},
		{
			name:       "",
			candidates: []int{1},
			target:     1,
			want:       [][]int{{1}},
		},
		{
			name:       "",
			candidates: []int{1},
			target:     2,
			want:       [][]int{{1, 1}},
		},
		{
			name:       "",
			candidates: []int{2, 7, 6, 3, 5, 1},
			target:     9,
			want:       [][]int{{1, 1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 2}, {1, 1, 1, 1, 1, 1, 3}, {1, 1, 1, 1, 1, 2, 2}, {1, 1, 1, 1, 2, 3}, {1, 1, 1, 1, 5}, {1, 1, 1, 2, 2, 2}, {1, 1, 1, 3, 3}, {1, 1, 1, 6}, {1, 1, 2, 2, 3}, {1, 1, 2, 5}, {1, 1, 7}, {1, 2, 2, 2, 2}, {1, 2, 3, 3}, {1, 2, 6}, {1, 3, 5}, {2, 2, 2, 3}, {2, 2, 5}, {2, 7}, {3, 3, 3}, {3, 6}},
		},
		{
			name:       "",
			candidates: []int{30, 34, 25, 24, 29, 38, 36, 42, 45, 44, 31, 28, 26, 37, 23, 20, 47, 40, 49, 46, 39, 43, 33, 41, 27, 32, 35, 48},
			target:     54,
			want:       [][]int{{20, 34}, {23, 31}, {24, 30}, {25, 29}, {26, 28}, {27, 27}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum(tt.candidates, tt.target)
			require.Equal(t, tt.want, got)
		})
	}
}
