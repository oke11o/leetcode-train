package _0xx

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
