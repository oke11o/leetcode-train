package _1xx

import (
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

/*
2154. Keep Multiplying Found Values by Two
https://leetcode.com/problems/keep-multiplying-found-values-by-two/description/
Easy
#array #sorting
*/
func findFinalValue(nums []int, original int) int {
	sort.Ints(nums)
	for _, v := range nums {
		if v == original {
			original *= 2
		}
	}

	return original
}

func Test_findFinalValue(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		original int
		want     int
	}{
		{
			name:     "",
			nums:     []int{5, 3, 6, 1, 12},
			original: 3,
			want:     24,
		},
		{
			name:     "",
			nums:     []int{2, 7, 9},
			original: 4,
			want:     4,
		},
		{
			name:     "",
			nums:     []int{2},
			original: 2,
			want:     4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findFinalValue(tt.nums, tt.original)
			require.Equal(t, tt.want, got)
		})
	}
}
