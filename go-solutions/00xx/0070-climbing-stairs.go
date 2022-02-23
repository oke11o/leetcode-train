package _0xx

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	prev := 1
	result := 1
	for i := 2; i <= n; i++ {
		t := result
		result = t + prev
		prev = t
	}
	return result
}
