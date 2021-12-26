package _0xx

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
 * 1014. Best Sightseeing Pair
 * Medium
 * #array, #dynamic_programming
 * https://leetcode.com/problems/best-sightseeing-pair/
 */
func maxScoreSightseeingPair(values []int) int {
	prev := values[0] + values[1] - 1
	res := prev
	for i := 2; i < len(values); i++ {
		prev = maxInt(
			values[i]+values[i-1]-1,
			prev-1+values[i]-values[i-1],
		)
		if res < prev {
			res = prev
		}
	}
	return res
}
