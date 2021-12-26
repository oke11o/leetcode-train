package _0xx

var resultPermute [][]int

/**
 * 0046. Permutations
 * Medium
 * #backtracking
 */
func permute(nums []int) [][]int {
	resultPermute = make([][]int, 0)
	permuteBacktrack(0, nums)
	return resultPermute
}

func permuteBacktrack(current int, nums []int) {
	if current == len(nums) {
		item := make([]int, len(nums))
		copy(item, nums)
		resultPermute = append(resultPermute, item)
		return
	}
	for i := current; i < len(nums); i++ {
		nums[i], nums[current] = nums[current], nums[i]
		permuteBacktrack(current+1, nums)
		nums[i], nums[current] = nums[current], nums[i]
	}
}
