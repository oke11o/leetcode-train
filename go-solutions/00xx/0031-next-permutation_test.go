package _0xx

import (
	"reflect"
	"testing"
)

/**
31. Next Permutation
Medium
https://leetcode.com/problems/next-permutation/

*/
func nextPermutation(nums []int) {
	idx := len(nums) - 2
	// Find the largest index idx such that nums[idx] < nums[idx + 1]. If no such index exists, just reverse nums and done.
	for idx >= 0 && nums[idx] > nums[idx+1] {
		idx--
	}
	if idx >= 0 { // may be 0, for example [2,8,6,3,2,1]. And
		// Find the largest index tmpx > idx such that nums[idx] >= nums[tmpx].
		tmpx := len(nums) - 1
		for nums[idx] >= nums[tmpx] {
			tmpx--
		}
		nums[tmpx], nums[idx] = nums[idx], nums[tmpx]
	}
	// idx == -1 when array is sorted desc
	reverse(nums[idx+1:])

}

func reverse(nums []int) {
	left := 0
	right := len(nums) - 1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

/*********************************/
/************* TESTS *************/
/*********************************/

func Test_nextPermutation(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "",
			nums: []int{1, 2, 3},
			want: []int{1, 3, 2},
		},
		{
			name: "",
			nums: []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
		{
			name: "",
			nums: []int{1, 1, 5},
			want: []int{1, 5, 1},
		},
		{
			name: "",
			nums: []int{1, 5, 8, 4, 7, 6, 5, 3, 1},
			want: []int{1, 5, 8, 5, 1, 3, 4, 6, 7},
		},
		{
			name: "",
			nums: []int{2, 8, 6, 3, 2, 1},
			want: []int{3, 1, 2, 2, 6, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.nums)
			if !reflect.DeepEqual(tt.want, tt.nums) {
				t.Errorf("nextPermutation: want=%+v; got=%+v", tt.want, tt.nums)
			}
		})
	}
}