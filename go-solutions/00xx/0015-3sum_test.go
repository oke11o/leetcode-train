package _0xx

import (
	"reflect"
	"sort"
	"testing"
)

/**
 * 0015. 3Sum
 * https://leetcode.com/problems/3sum/
 * Medium
 *
 * [#array, #two_pointers, #sorting]
 *
 */
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		target := nums[i]
		j := i + 1
		k := len(nums) - 1
		for j < k || -target != nums[j]+nums[k] {
			if nums[j]+nums[k] > 0-target {
				k--
			} else {
				j++
			}
		}
		if -target == nums[j]+nums[k] {
			result = append(result, []int{nums[i], nums[j], nums[k]})
		}
	}
	return result
}

func Test_threeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "",
			nums: []int{},
			want: [][]int{},
		},
		{
			name: "",
			nums: []int{0},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
