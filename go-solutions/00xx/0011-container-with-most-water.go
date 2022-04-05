package _0xx

/**
https://leetcode.com/problems/container-with-most-water/
11. Container With Most Water
Medium
#two_pointer
*/
func maxArea(height []int) int {
	result := 0
	l := 0
	r := len(height) - 1
	for l < r {
		h := height[l]
		w := r - l
		if h > height[r] {
			h = height[r]
			r--
		} else {
			l++
		}
		area := h * w
		if result < area {
			result = area
		}
	}

	return result
}
