package _0xx

/**
 * 0033. Search in Rotated Sorted Array
 * Medium
 * #array, #binary_search
 * https://leetcode.com/problems/search-in-rotated-sorted-array/
 *
 * Reference: 0153. Find Minimum in Rotated Sorted Array
 *
 * UNSOLVED, TODO
 */
func search(nums []int, target int) int {
	left := -1
	right := len(nums) - 1
	for (right - left) > 1 {
		i := (right + left) / 2

		if nums[i] > nums[left] {

		}

		if nums[i] > target {
			right = i
		} else {
			left = i
		}
	}
	return 0
}
