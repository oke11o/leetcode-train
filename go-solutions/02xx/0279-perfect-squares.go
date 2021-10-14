package _2xx

const MaxInt32 = 1<<31 - 1

var cache = make(map[int]int)

// 0279. Perfect Squares (ver 2)
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ { // dp[0] == 0
		dp[i] = n + 1
	}

	for target := 1; target < n+1; target++ {
		for s := 1; s < target+1; s++ {
			square := s * s
			if target-square < 0 {
				break
			}
			if dp[target] > dp[target-square]+1 {
				dp[target] = dp[target-square] + 1
			}
		}
	}

	return dp[n]
}

// 0279. Perfect Squares
func numSquares_(n int) int {
	if n < 4 {
		return n
	}
	if v, ok := cache[n]; ok {
		return v
	}
	mins := []int{}
	for i := 1; i*i <= n; i++ {
		ii := i * i
		if ii == n {
			return 1
		}
		mins = append(mins, 1+numSquares(n-ii))
	}

	res := min(mins...)
	cache[n] = res
	return res
}

func min(nums ...int) int {
	res := int(MaxInt32)
	for _, v := range nums {
		if res > v {
			res = v
		}
	}
	return res
}
