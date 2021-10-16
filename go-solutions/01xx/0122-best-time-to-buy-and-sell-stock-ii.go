package _1xx

// 0122. Best Time to Buy and Sell Stock II
func maxProfit_ii(prices []int) int {
	result := 0
	for i := 1; i < len(prices); i++ {
		if amount := prices[i] - prices[i-1]; amount > 0 {
			result += amount
		}
	}
	return result
}
