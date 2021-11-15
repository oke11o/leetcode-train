package _1xx

// 0191. Number of 1 Bits
func hammingWeight(num uint32) int {
	result := 0
	for num != 0 {
		result += int(num % 2)
		num >>= 1
	}
	return result
}
