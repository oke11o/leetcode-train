package _0xx

import "testing"

/**
https://leetcode.com/problems/missing-element-in-sorted-array/
1060. Missing Element in Sorted Array
Medium
*/
func missingElement(nums []int, k int) int {
	i := nums[0]
	numsIdx := 1
	if len(nums) > 1 {
		for i = i + 1; i < nums[len(nums)-1]; i++ {
			//check is missing number
			if i == nums[numsIdx] {
				numsIdx++
				continue
			}
			k--
			if k == 0 {
				return i
			}
		}
	}

	return i + k
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_missingElement(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{4, 7, 9, 10},
				k:    1,
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				nums: []int{4, 7, 9, 10},
				k:    3,
			},
			want: 8,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 4},
				k:    3,
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				nums: []int{1},
				k:    1,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				k:    1,
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				k:    4,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingElement(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("missingElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
