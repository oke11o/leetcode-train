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
 *
 * Next challenges
 * - https://leetcode.com/problems/flatten-2d-vector/
 * - https://leetcode.com/problems/shortest-completing-word/
 * - https://leetcode.com/problems/unique-email-addresses/
 */
func search81(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	endIdx := len(nums) - 1
	state := 0 // 0 - undedined, 1 - inLeftPart, 2 - inRightPart

	var startIdx int
	if nums[endIdx] == nums[startIdx] {
		state = 0
	} else if target < nums[startIdx] {
		state = 2
	} else {
		state = 1
	}

	left := -1
	right := len(nums)
	for (right - left) > 1 {
		i := (right + left) / 2
		v := nums[i]
		if v == target {
			return true
		}
		if state == 0 {
			if nums[endIdx] == nums[startIdx] {
				left++
				startIdx = left
			} else {
				if target < nums[startIdx] {
					state = 2
				} else {
					state = 1
				}
			}
			continue
		}
		if state == 1 {
			if v < nums[startIdx] {
				right = i
			} else if v > target {
				right = i
			} else {
				left = i
			}
		} else {
			if v > nums[endIdx] {
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
	return left != -1
}
