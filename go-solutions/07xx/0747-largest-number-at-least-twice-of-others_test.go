package _7xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

/*
747. Largest Number At Least Twice of Others
https://leetcode.com/problems/largest-number-at-least-twice-of-others/
Easy
#array #sorting
*/
func dominantIndex(nums []int) int {
	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[idx] < nums[i] {
			idx = i
		}
	}
	for i := 0; i < len(nums); i++ {
		if i != idx && nums[i]*2 > nums[idx] {
			return -1
		}
	}
	return idx
}

func Test_dominantIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "",
			nums: []int{3, 6, 1, 0},
			want: 1,
		},
		{
			name: "",
			nums: []int{1, 2, 3, 4},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dominantIndex(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
