package _1xx

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	minus3 := 0
	minus2 := 0
	minus1 := 1
	result := 0
	for i := 2; i <= n; i++ {
		result = minus3 + minus2 + minus1
		minus3, minus2, minus1 = minus2, minus1, result
	}
	return result
}
