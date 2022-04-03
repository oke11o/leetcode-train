package _1xx

import "sort"

/**
https://leetcode.com/problems/minimum-sum-of-four-digit-number-after-splitting-digits/
2160. Minimum Sum of Four Digit Number After Splitting Digits
Easy
*/
func minimumSum(num int) int {
	digits := make([]int, 0, 4)
	for num > 9 {
		digits = append(digits, num%10)
		num = num / 10
	}
	digits = append(digits, num)
	sort.Ints(digits)

	return 10*(digits[0]+digits[1]) + digits[2] + digits[3]
}
