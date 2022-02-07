package _0xx

import (
	"reflect"
	"sort"
	"testing"
)

/**
0015. 3Sum
https://leetcode.com/problems/3sum/
Medium

[#array, #two_pointers, #sorting]
#TODO: wip
*/
func threeSum(nums []int) [][]int {
	var twoSum_twoPointers = func(numbers []int, curIdx int) []int {
		target := 0 - numbers[curIdx]
		right := len(numbers) - 1
		left := curIdx + 1
		for left < right {
			sum := numbers[left] + numbers[right]
			if sum == target {
				return []int{numbers[curIdx], numbers[left], numbers[right]}
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
		return nil
	}
	_ = twoSum_twoPointers

	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if res := twoSum_twoPointers(nums, i); res != nil {
			result = append(result, res)
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
			nums: []int{-1, 0, 1, 2, -1, -4}, // -4,-1,-1,0,1,2
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
		{
			name: "",
			nums: []int{0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "",
			nums: []int{-4, -1, -1, -1, 0, 1, 2, 3}, // -4,-1,-1,0,1,2
			want: [][]int{{-4, 1, 3}, {-1, -1, 2}, {-1, 0, 1}},
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
