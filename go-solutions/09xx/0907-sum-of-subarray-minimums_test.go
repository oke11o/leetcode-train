package _9xx

import "testing"

/**
https://leetcode.com/problems/sum-of-subarray-minimums/
907. Sum of Subarray Minimums
Medium
#array, #dynamic, #stack, #amazon
*/
func sumSubarrayMins(arr []int) int {
	return 0
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_sumSubarrayMins(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumSubarrayMins(tt.args.arr); got != tt.want {
				t.Errorf("sumSubarrayMins() = %v, want %v", got, tt.want)
			}
		})
	}
}
