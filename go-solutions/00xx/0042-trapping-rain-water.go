package _0xx

/**
https://leetcode.com/problems/trapping-rain-water/
42. Trapping Rain Water
Hard
#two_pointers
*/
func trap(height []int) int {
	left := 0
	right := len(height) - 1

	ans := 0

	leftMax := 0
	rightMax := 0

	for left < right {
		if height[left] < height[right] {
			if leftMax <= height[left] {
				leftMax = height[left]
			} else {
				ans += leftMax - height[left]
			}
			left++
		} else {
			if rightMax <= height[right] {
				rightMax = height[right]
			} else {
				ans += rightMax - height[right]
			}
			right--
		}
	}

	return ans
}
