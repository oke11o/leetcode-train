package _0xx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
https://leetcode.com/problems/largest-rectangle-in-histogram/
84. Largest Rectangle in Histogram
Hard

How we calculate rectangle? Multiply the width by the height
On each bar we should know.
Current height.
Count of left bars for first bar with smaller height.
Count of right bars for first bar with smaller height.
And now we can calculate rectangle. And compare with max.

Ok. How we can calculate left and right bars?
Oh.

*/

/**
Divide and Conquer Approach
*/
func largestRectangleArea(heights []int) int {
	const maxInt32 = 1<<31 - 1
	var calculateArea func(start, end int) int
	calculateArea = func(start, end int) int {
		if start > end {
			return 0
		}
		minIndex := start
		for i := start; i <= end; i++ {
			if heights[minIndex] > heights[i] {
				minIndex = i
			}
		}
		curArea := heights[minIndex] * (end - start + 1)
		leftArea := calculateArea(start, minIndex-1)
		if curArea < leftArea {
			curArea = leftArea
		}
		rightArea := calculateArea(minIndex+1, end)
		if curArea < rightArea {
			curArea = rightArea
		}
		return curArea
	}
	return calculateArea(0, len(heights)-1)
}

/**
Brute force
*/
func largestRectangleArea_1(heights []int) int {
	const maxInt32 = 1<<31 - 1
	result := 0
	for i := 0; i < len(heights); i++ {
		minHeight := maxInt32
		for j := i; j < len(heights); j++ {
			if minHeight > heights[j] {
				minHeight = heights[j]
			}
			area := minHeight * (j - i + 1)
			if result < area {
				result = area
			}
		}
	}

	return result
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
		{
			name: "",
			args: args{
				heights: []int{2, 1, 5, 6, 2, 3},
			},
			want: 10,
		},
		{
			name: "",
			args: args{
				heights: []int{2, 4},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, largestRectangleArea(tt.args.heights), "largestRectangleArea(%v)", tt.args.heights)
		})
	}
}
