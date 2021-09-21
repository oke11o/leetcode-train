package _5xx

// 509. Fibonacci Number
// F(0) = 0, F(1) = 1
// F(n) = F(n - 1) + F(n - 2), for n > 1.
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	minus2 := 0
	minus1 := 1
	result := 0
	for i := 2; i <= n; i++ {
		result = minus2 + minus1
		minus2, minus1 = minus1, result
	}
	return result
}
