package _1xx

/*
121. Best Time to Buy and Sell Stock
[7,1,5,3,6,4]
*/
func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	result := 0
	min := prices[0]
	for _, val := range prices {
		newMax := val - min
		if result < newMax {
			result = newMax
		}
		if min > val {
			min = val
		}
	}
	return result
}

func maxProfit(prices []int) int {
	dp := make([]int, len(prices)+1)
	for i := 1; i < len(prices); i++ {

	}
	return dp[len(dp)-1]
}
