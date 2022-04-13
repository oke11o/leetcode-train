package _0xx

/**
https://leetcode.com/problems/spiral-matrix-ii/
59. Spiral Matrix II
Medium
*/
func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	cnt := 1
	for layer := 0; layer < (n+1)/2; layer++ {
		// direction 1 - traverse from left to right
		for ptr := layer; ptr < n-layer; ptr++ {
			if result[layer] == nil {
				result[layer] = make([]int, n)
			}
			result[layer][ptr] = cnt
			cnt++
		}
		// direction 2 - traverse from top to bottom
		for ptr := layer + 1; ptr < n-layer; ptr++ {
			if result[ptr] == nil {
				result[ptr] = make([]int, n)
			}
			result[ptr][n-layer-1] = cnt
			cnt++
		}
		// direction 3 - traverse from right to left
		for ptr := n - layer - 2; ptr >= layer; ptr-- {
			if result[n-layer-1] == nil {
				result[n-layer-1] = make([]int, n)
			}
			result[n-layer-1][ptr] = cnt
			cnt++
		}
		// direction 4 - traverse from bottom to top
		for ptr := n - layer - 2; ptr > layer; ptr-- {
			if result[ptr] == nil {
				result[ptr] = make([]int, n)
			}
			result[ptr][layer] = cnt
			cnt++
		}
	}

	return result
}
