package go_solutions

func rob(nums []int) int {
	var max = func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	dp := make([]int, len(nums)+2)
	for i, cur := range nums {
		dp[i+2] = max(dp[i+1], dp[i]+cur)
	}
	return dp[len(dp)-1]
}
