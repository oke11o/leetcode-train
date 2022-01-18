package _2xx

/**
 * 1283. Find the Smallest Divisor Given a Threshold
 * https://leetcode.com/problems/find-the-smallest-divisor-given-a-threshold/
 * Medium
 *
 * Array, Binary Search
 *
 * My answers:
 *  - what max length of input array?
 *  - is it posible zero value on any elem? (doesn't matter) 0/3 = 0
 *  - is it possible when input array is empty?
 *  - is input array is asc ordered? Or not?
 * What corner cases?
 * - I don't know
 *
 * 4. Brute force solution
 */
func smallestDivisor(nums []int, threshold int) int {
	left := 0
	right := 0 // largest element
	for _, v := range nums {
		if v > right {
			right = v
		}
	}
	for right-left > 1 { // 1-0 == 1 and it if out of out condition
		divisor := (right + left) / 2
		sum := smallestDivisor_dividedListSum(nums, divisor)
		if sum > threshold { // 7>6 = true, what is mean? Is it mean divisor=4 is wrong. And
			left = divisor
		} else {
			right = divisor
		}
	}

	return right
}

func smallestDivisor_dividedListSum(in []int, divisor int) int {
	sum := 0
	for _, i := range in {
		sum += i / divisor
		if i%divisor > 0 {
			sum++
		}
	}
	return sum
}
