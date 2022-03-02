package _0xx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
80. Remove Duplicates from Sorted Array II
Medium

Similar - https://leetcode.com/problems/remove-duplicates-from-sorted-array/
*/
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	flag := false
	pos := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[pos] = nums[i]
			pos++
			flag = false
		} else if !flag {
			nums[pos] = nums[i]
			pos++
			flag = true
		}
	}
	return pos
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_removeDuplicates(t *testing.T) {
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
				nums: []int{1, 1, 1, 2, 2, 3},
			},
			want: []int{1, 1, 2, 2, 3},
		},
		{
			name: "",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			},
			want: []int{0, 0, 1, 1, 2, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args.nums)
			assert.Equalf(t, len(tt.want), got, "removeDuplicates(%v)", tt.args.nums)
			assert.Equalf(t, tt.want, tt.args.nums[0:got], "removeDuplicates(%v)", tt.args.nums)
		})
	}
}
