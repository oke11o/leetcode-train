package go_solutions

func permute(nums []int) [][]int {
	length := len(nums)
	if length == 1 {
		return [][]int{{nums[0]}}
	}
	result := make([][]int, 0, length^2)
	for i := 0; i < length; i++ {
		for j := length - 1; j > 0; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
			item := make([]int, length)
			copy(item, nums)
			result = append(result, item)
		}
	}
	return result
}
