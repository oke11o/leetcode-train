package _0xx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
https://leetcode.com/problems/largest-rectangle-in-histogram/
84. Largest Rectangle in Histogram
Hard
*/
func largestRectangleArea(heights []int) int {
	return 0
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_largestRectangleArea(t *testing.T) {
	type args struct {
		heights []int
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
			assert.Equalf(t, tt.want, largestRectangleArea(tt.args.heights), "largestRectangleArea(%v)", tt.args.heights)
		})
	}
}
