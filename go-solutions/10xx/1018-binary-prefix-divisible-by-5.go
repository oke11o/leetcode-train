package _0xx

/**
 * 1018. Binary Prefix Divisible By 5
 * Easy
 * #array
 * https://leetcode.com/problems/binary-prefix-divisible-by-5/
 */
func prefixesDivBy5(nums []int) []bool {
	result := make([]bool, len(nums))
	var t int
	for i, n := range nums {
		t = t<<1 + n
		t %= 5
		if t == 0 {
			result[i] = true
		}
	}

	return result
}
