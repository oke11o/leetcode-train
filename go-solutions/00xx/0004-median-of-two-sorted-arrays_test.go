package _0xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// 0004. Median of Two Sorted Arrays
// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
// The overall run time complexity should be O(log (m+n)).
//
// https://leetcode.com/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 1. Length of result list

	resLength := len(nums1) + len(nums2)
	p1 := 0
	p2 := 0
	cur := 0
	prev := 0
	for p1+p2 < resLength/2+1 {
		// if nums1[p1] {}
	}
	if resLength%2 != 0 {
		return float64(prev)
	}

	return float64(prev+cur) / 2
}

// 0,1,2,3,4
// 	   ^ ^
// 0,1,2,3
//   ^ ^

func Test_findMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			name:  "",
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2.0,
		},
		{
			name:  "",
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2.0,
		},
		{
			name:  "",
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMedianSortedArrays(tt.nums1, tt.nums2)
			require.Equalf(t, tt.want, result, "findMedianSortedArrays(%v, %v)", tt.nums1, tt.nums2)
		})
	}
}
