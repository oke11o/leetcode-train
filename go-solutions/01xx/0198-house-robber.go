package _1xx

func rob(nums []int) int {
	dp := make([]int, len(nums)+2)
	for i, num := range nums {
		dp[i+2] = dp[i] + num
		if dp[i+2] < dp[i+1] {
			dp[i+2] = dp[i+1]
		}
	}
	return dp[len(dp)-1]
}
