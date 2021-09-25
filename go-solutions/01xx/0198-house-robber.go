package _1xx

func rob(nums []int) int {
	var max = func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	dp := make([]int, len(nums)+2)
	for i := 2; i < len(dp); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-2])
	}
	return dp[len(dp)-1]
}
