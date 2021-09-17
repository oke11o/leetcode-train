package go_solutions

import (
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	nums1 = append(nums1, 9999)
	nums2 = append(nums2, 9999)
	result := make([]int, 0, len(nums1))
	var i1, i2 int
	for i1 < len(nums1) && i2 < len(nums2) {
		if nums1[i1] == nums2[i2] {
			result = append(result, nums1[i1])
			i1++
			i2++
		} else if nums1[i1] < nums2[i2] {
			i1++
		} else {
			i2++
		}
	}

	return result[:len(result)-1]
}
