package _4xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// 0485. Max Consecutive Ones
// https://leetcode.com/problems/max-consecutive-ones
func findMaxConsecutiveOnes(nums []int) int {
	result := 0
	current := 0
	for _, n := range nums {
		if n != 0 {
			current++
			continue
		}
		if result < current {
			result = current
		}
		current = 0
	}
	if result < current {
		result = current
	}
	current = 0
	return result
}

func Test_findMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "",
			nums: []int{1, 1, 0, 1, 1, 1},
			want: 3,
		},
		{
			name: "",
			nums: []int{1, 0, 1, 1, 0, 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxConsecutiveOnes(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
