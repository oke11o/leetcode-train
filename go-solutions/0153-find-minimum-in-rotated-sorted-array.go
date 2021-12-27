package go_solutions

/**
 * 0153. Find Minimum in Rotated Sorted Array
 * Medium
 * #array, #binary_search
 * #amazon
 * https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
 * References: 33. Search in Rotated Sorted Array
 */
func findMinInRotatedSortedArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left := 0
	right := len(nums) - 1
	if nums[left] < nums[right] {
		return nums[0]
	}
	for (right - left) > 1 {
		i := (left + right) / 2
		if nums[i] > nums[left] {
			left = i
		} else {
			right = i
		}
	}
	return nums[left+1]
}
