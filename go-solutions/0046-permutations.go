package go_solutions

var result [][]int

func permute(nums []int) [][]int {
	result = make([][]int, 0)
	backtrack(0, nums)
	return result
}

func backtrack(current int, nums []int) {
	if current == len(nums) {
		item := make([]int, len(nums))
		copy(item, nums)
		result = append(result, item)
		return
	}
	for i := current; i < len(nums); i++ {
		nums[i], nums[current] = nums[current], nums[i]
		backtrack(current+1, nums)
		nums[i], nums[current] = nums[current], nums[i]
	}
}
