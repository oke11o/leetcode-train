package _0xx

// 0053. Maximum Subarray
func maxSubArray(nums []int) int {
	result := nums[0]
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]

		dp[i] = num
		candidateSum := dp[i-1] + num
		if candidateSum > num {
			dp[i] = candidateSum
		}

		if result < dp[i] {
			result = dp[i]
		}
	}
	return result
}
