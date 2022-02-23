package _4xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/**
https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
448. Find All Numbers Disappeared in an Array
Easy
*/
func findDisappearedNumbers(nums []int) []int {
	var result []int
	i := 0
	for i < len(nums) {
		pos := nums[i] - 1
		if nums[i] != nums[pos] {
			nums[i], nums[pos] = nums[pos], nums[i]
		} else {
			i++
		}
	}
	for i, n := range nums {
		if i != n-1 {
			result = append(result, i+1)
		}
	}
	return result
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_findDisappearedNumbers(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "",
			nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			want: []int{5, 6},
		},
		{
			name: "",
			nums: []int{1, 1},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDisappearedNumbers(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
