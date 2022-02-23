package _4xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/**
https://leetcode.com/problems/find-all-duplicates-in-an-array/
442. Find All Duplicates in an Array
Medium
*/
func findDuplicates(nums []int) []int {
	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		newKey := nums[i]
		if newKey < 0 {
			newKey = 0 - newKey
		}
		newKey--
		if nums[newKey] < 0 {
			result = append(result, newKey+1)
		} else {
			nums[newKey] = 0 - nums[newKey]
		}
	}
	return result
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_findDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "",
			nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			want: []int{2, 3},
		},
		{
			name: "",
			nums: []int{1, 1, 2},
			want: []int{1},
		},
		{
			name: "",
			nums: []int{1},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDuplicates(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
