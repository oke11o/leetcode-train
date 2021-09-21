package _4xx

// 0485. Max Consecutive Ones
func findMaxConsecutiveOnes(nums []int) int {
	result := 0
	current := 0
	for _, n := range nums {
		if n != 0 {
			current++
			continue
		}
		if result < current {
			result = current
		}
		current = 0
	}
	if result < current {
		result = current
	}
	current = 0
	return result
}
