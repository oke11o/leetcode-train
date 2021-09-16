package go_solutions

func maxProfit(prices []int) int {
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
