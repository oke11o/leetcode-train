package _4xx

import (
	"reflect"
	"testing"
)

/**
1480. Running Sum of 1d Array
Easy
https://leetcode.com/problems/running-sum-of-1d-array/
*/
func runningSum(nums []int) []int {
	result := make([]int, len(nums))
	result[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] + nums[i]
	}
	return result
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_runningSum(t *testing.T) {
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
				nums: []int{1, 2, 3, 4},
			},
			want: []int{1, 3, 6, 10},
		},
		{
			name: "",
			args: args{
				nums: []int{1, 1, 1, 1, 1},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "",
			args: args{
				nums: []int{3, 1, 2, 10, 1},
			},
			want: []int{3, 4, 6, 16, 17},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runningSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runningSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
