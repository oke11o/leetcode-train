package go_solutions

func findDisappearedNumbers(nums []int) []int {
	var result []int
	i := 0
	for i < len(nums) {
		pos := nums[i] - 1
		if nums[i] != nums[pos] {
			nums[i], nums[pos] = nums[pos], nums[i]
		} else {
			i++
		}
	}
	for i, n := range nums {
		if i != n-1 {
			result = append(result, i+1)
		}
	}
	return result
}
