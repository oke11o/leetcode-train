package _2xx

/**
 * 209. Minimum Size Subarray Sum
 * https://leetcode.com/problems/minimum-size-subarray-sum/
 * Medium
 * [Sliding Window, #sliding_window, #array]
 *
 * Questions:
 * - Can be nums be empty?
 * - Is it possible, we can't find result (result == 0) ?
 */
func minSubArrayLen(target int, nums []int) int {
	result := len(nums) + 1
	var windowSum, leftP int
	for rightP, n := range nums {
		windowSum += n
		for windowSum >= target {
			if result > rightP-leftP+1 {
				result = rightP - leftP + 1 // 15
			}
			windowSum -= nums[leftP]
			leftP++
		}
	}
	if result == len(nums)+1 {
		return 0
	}
	return result
}
