package _0xx

/**
 * 0033. Search in Rotated Sorted Array
 * Medium
 * #array, #binary_search
 * https://leetcode.com/problems/search-in-rotated-sorted-array/
 *
 * Reference:
 * - 0153. Find Minimum in Rotated Sorted Array
 * - 0081. Search in Rotated Sorted Array II
 */
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	lastIdx := len(nums) - 1
	inFirstPart := true
	if target < nums[0] {
		inFirstPart = false
	}

	left := -1
	right := len(nums)
	for (right - left) > 1 {
		i := (right + left) / 2
		if inFirstPart {
			if nums[i] < nums[0] {
				right = i
			} else if nums[i] > target {
				right = i
			} else {
				left = i
			}
		} else {
			if nums[i] > nums[lastIdx] {
				left = i
			} else if nums[i] > target {
				right = i
			} else {
				left = i
			}
		}
	}
	if left != -1 && nums[left] != target {
		return -1
	}
	return left
}
