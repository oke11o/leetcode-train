package _3xx

func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	result := make([]int, n+1)
	result[1] = 1
	for i := 2; i <= n; i++ {
		even := i % 2
		result[i] = result[i/2] + even
	}

	return result
}
