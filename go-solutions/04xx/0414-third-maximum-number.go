package _4xx

/**
 * 0414. Third Maximum Number
 * Easy
 * #array, #sorting
 * https://leetcode.com/problems/third-maximum-number/
 */
func thirdMax(nums []int) int {
	maxMin := -1 << (64 - 1)
	m1 := maxMin
	m2 := m1
	m3 := m2

	for _, n := range nums {
		if m1 == n {
			continue
		}
		if m1 < n {
			m3 = m2
			m2 = m1
			m1 = n
		} else if m2 < n && m1 != n {
			m3 = m2
			m2 = n
		} else if m3 < n && m2 != n {
			m3 = n
		}
	}
	if m3 == maxMin {
		return m1
	}
	return m3
}
