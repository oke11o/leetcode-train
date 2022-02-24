package _0xx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
https://leetcode.com/problems/first-missing-positive/
41. First Missing Positive
Hard
*/
//3, 4, -1, 1
//-1,4,3,1
//++
//-1,1,3,4
//1,-1,3,4
//++
//done
func firstMissingPositive(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] > 0 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{0, 1, 2},
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				nums: []int{3, 4, -1, 1},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums: []int{7, 8, 9, 11, 12},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				nums: []int{1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, firstMissingPositive(tt.args.nums), "firstMissingPositive(%v)", tt.args.nums)
		})
	}
}
