package _1xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/**
 * 0153. Find Minimum in Rotated Sorted Array
 * Medium
 * #array, #binary_search
 * #amazon
 * https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
 * References: 33. Search in Rotated Sorted Array
 */
func findMinInRotatedSortedArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left := 0
	right := len(nums) - 1
	if nums[left] < nums[right] {
		return nums[0]
	}
	for (right - left) > 1 {
		i := (left + right) / 2
		if nums[i] > nums[left] {
			left = i
		} else {
			right = i
		}
	}
	return nums[left+1]
}

func Test_findMinInRotatedSortedArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "",
			nums: []int{3, 4, 5, 1, 2},
			want: 1,
		},
		{
			name: "",
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			want: 0,
		},
		{
			name: "",
			nums: []int{11, 13, 15, 17},
			want: 11,
		},
		{
			name: "",
			nums: []int{44},
			want: 44,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMinInRotatedSortedArray(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
