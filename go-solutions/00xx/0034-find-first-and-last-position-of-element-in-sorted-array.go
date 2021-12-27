package _0xx

/**
 * 34. Find First and Last Position of Element in Sorted Array
 * Medium
 * #array, #binary_search
 * https://www.educative.io/m/find-low-high-index
 * https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
 */
func searchRange(nums []int, target int) []int {
	left := -1
	right := len(nums)
	for (right - left) > 1 {
		i := (right + left) / 2
		if nums[i] > target {
			right = i
		} else {
			left = i
		}
	}
	if left == -1 || nums[left] != target {
		return []int{-1, -1}
	}

	// find left
	target--
	rightRes := left
	left = -1
	right = rightRes
	for right-left > 1 {
		i := (right + left) / 2
		if nums[i] > target {
			right = i
		} else {
			left = i
		}
	}

	return []int{left + 1, rightRes}
}
