package _9xx

// 0977. Squares of a Sorted Array
func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	left := 0
	right := len(nums) - 1

	for i := right; i >= 0; i-- {
		if nums[left] < 0 && -nums[left] > nums[right] {
			result[i] = nums[left] * nums[left]
			left++
		} else {
			result[i] = nums[right] * nums[right]
			right--
		}
	}
	return result
}
