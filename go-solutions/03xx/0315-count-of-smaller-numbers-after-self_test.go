package _3xx

import (
	"reflect"
	"testing"
)

/**
https://leetcode.com/problems/count-of-smaller-numbers-after-self/
315. Count of Smaller Numbers After Self
Hard

Constraints:
1 <= nums.length <= 105
-104 <= nums[i] <= 104
//Approach 1: Segment Tree
*/
func countSmaller(nums []int) []int {
	offset := 10 ^ 4     // offset negative to non-negative
	size := 2*offset + 1 // total possible values in nums
	tree := make([]int, 2*size)
	_ = tree
	result := make([]int, len(nums))
	return result
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_countSmaller(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				nums: []int{5, 2, 6, 1},
			},
			want: []int{2, 1, 1, 0},
		},
		{
			name: "",
			args: args{
				nums: []int{2, 3, 4, 5, 2, 2, 4, 6, 8, 6, 4, 9, 5},
			},
			want: []int{0, 2, 2, 4, 0, 0, 0, 2, 3, 2, 0, 1, 0},
		},
		{
			name: "",
			args: args{
				nums: []int{-1},
			},
			want: []int{0},
		},
		{
			name: "",
			args: args{
				nums: []int{-1, -1},
			},
			want: []int{0, 0},
		},
		{
			name: "",
			args: args{
				nums: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSmaller(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSmaller() = %v, want %v", got, tt.want)
			}
		})
	}
}
