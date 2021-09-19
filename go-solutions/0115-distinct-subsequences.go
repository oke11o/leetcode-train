package go_solutions

// 0115. Distinct Subsequences
func numDistinct(str string, target string) int {
	strLen := len(str)
	tarLen := len(target)
	if strLen == tarLen {
		if str == target {
			return 1
		} else {
			return 0
		}
	} else if strLen < tarLen {
		return 0
	}

	dp := make([][]int, tarLen+1)
	dp[0] = make([]int, strLen+1)
	for j := 0; j <= strLen; j++ {
		dp[0][j] = 1
	}

	for i := 0; i < tarLen; i++ {
		for j := 0; j < strLen; j++ {
			if dp[i+1] == nil {
				dp[i+1] = make([]int, strLen+1)
			}

			if str[j] == target[i] {
				dp[i+1][j+1] = dp[i][j] + dp[i+1][j]
			} else {
				dp[i+1][j+1] = dp[i+1][j]
			}
		}
	}

	return dp[tarLen][strLen]
}
