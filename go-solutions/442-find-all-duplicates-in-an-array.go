package go_solutions

func findDuplicates(nums []int) []int {
	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		newKey := nums[i]
		if newKey < 0 {
			newKey = 0 - newKey
		}
		newKey--
		if nums[newKey] < 0 {
			result = append(result, newKey+1)
		} else {
			nums[newKey] = 0 - nums[newKey]
		}
	}
	return result
}
