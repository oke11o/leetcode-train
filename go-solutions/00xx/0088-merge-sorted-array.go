package _0xx

// 0088. Merge Sorted Array
// Easy
// Array, Two Pointer, Sorting
// https://leetcode.com/problems/merge-sorted-array/
func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	idx := m + n - 1
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[idx] = nums1[m-1]
			m--
		} else {
			nums1[idx] = nums2[n-1]
			n--
		}
		idx--
	}
	for n > 0 {
		nums1[idx] = nums2[n-1]
		n--
		idx--
	}
}
