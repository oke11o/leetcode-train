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
	windowLen := 0
	windowSum := 0
	leftP := 0
	for rightP := 0; rightP < len(nums); rightP++ {
		windowLen++
		windowSum += nums[rightP]
		for windowSum > target {
			windowSum -= nums[leftP]
			leftP++
			windowLen--
		}
		if windowSum == target && result > windowLen {
			result = windowLen
		}
	}
	if result > len(nums) {
		return 0
	}
	return result
}
