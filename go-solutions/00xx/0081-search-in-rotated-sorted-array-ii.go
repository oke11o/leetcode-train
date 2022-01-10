package _0xx

/**
 * 81. Search in Rotated Sorted Array II
 * Medium
 * #array, #binary_search
 * https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
 *
 * Reference:
 * - 0153. Find Minimum in Rotated Sorted Array
 * - 0033. Search in Rotated Sorted Array

 * UNSOLVED, TODO
 */
func search81(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
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
		v := nums[i]
		if v == target {
			return true
		}
		tmpLeft := left
		if tmpLeft == -1 {
			tmpLeft = 0
		}
		if v == nums[tmpLeft] {
			left++
			continue
		}
		if inFirstPart {
			if v < nums[0] {
				right = i
			} else if v > target {
				right = i
			} else {
				left = i
			}
		} else {
			if v > nums[lastIdx] {
				left = i
			} else if v > target {
				right = i
			} else {
				left = i
			}
		}
	}
	if left != -1 && nums[left] != target {
		return false
	}
	return true
}
