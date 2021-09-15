package go_solutions

func sumSquareDigits(n int) int {
	result := 0
	for n != 0 {
		v := n % 10
		result += v * v
		n /= 10
	}
	return result
}

func isHappy(n int) bool {
	dict := make(map[int]struct{})
	for {
		dict[n] = struct{}{}
		n = sumSquareDigits(n)
		_, ok := dict[n]
		if n == 1 || ok {
			break
		}
	}
	return n == 1
}
