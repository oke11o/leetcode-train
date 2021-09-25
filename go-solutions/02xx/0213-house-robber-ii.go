package _2xx

// 0213. House Robber II
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	v1 := rob_s(nums[1:])
	v2 := rob_s(nums[:len(nums)-1])
	if v1 > v2 {
		return v1
	}

	return v2
}

func rob_s(nums []int) int {
	dp := make([]int, len(nums)+2)
	for i, num := range nums {
		dp[i+2] = dp[i] + num
		if dp[i+2] < dp[i+1] {
			dp[i+2] = dp[i+1]
		}
	}
	return dp[len(dp)-1]
}
