package go_solutions

var resultPermute [][]int

func permute(nums []int) [][]int {
	resultPermute = make([][]int, 0)
	backtrack(0, nums)
	return resultPermute
}

func backtrack(current int, nums []int) {
	if current == len(nums) {
		item := make([]int, len(nums))
		copy(item, nums)
		resultPermute = append(resultPermute, item)
		return
	}
	for i := current; i < len(nums); i++ {
		nums[i], nums[current] = nums[current], nums[i]
		backtrack(current+1, nums)
		nums[i], nums[current] = nums[current], nums[i]
	}
}
