package _0xx

// 0053. Maximum Subarray
// Same as Kadane. But
func maxSubArray(nums []int) int {
	result := nums[0]
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]

		dp[i] = num
		if dp[i-1]+num > num { // Если к предыдущ результату добавить текущий шаг, будет ли новый результат лучше, чем просто текущий шаг?
			dp[i] = dp[i-1] + num
		}

		if result < dp[i] {
			result = dp[i]
		}
	}
	return result
}

// Kadane's algorithm
func kadane(nums []int) int {
	result := nums[0]
	curSum := nums[0]
	for i := 1; i < len(nums); i++ {
		curSum += nums[i]
		if curSum < nums[i] { // Если новая сумма хуже, чем просто данное значение, используем просто значение
			curSum = nums[i]
		}
		if result < curSum {
			result = curSum
		}
	}
	return result
}
