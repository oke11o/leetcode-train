package go_solutions

import "strconv"

var resultPermuteUnique [][]int

func permuteUnique(nums []int) [][]int {
	dubles := make(map[string]struct{})
	resultPermuteUnique = make([][]int, 0)
	backtrackPermuteUnieque(0, nums, dubles)
	return resultPermuteUnique
}

func backtrackPermuteUnieque(current int, nums []int, doubles map[string]struct{}) {
	if current == len(nums) {
		var key string
		for _, j := range nums {
			key += strconv.Itoa(j)
		}
		if _, ok := doubles[key]; ok {
			return
		}
		doubles[key] = struct{}{}
		item := make([]int, len(nums))
		copy(item, nums)
		resultPermuteUnique = append(resultPermuteUnique, item)
		return
	}

	for i := current; i < len(nums); i++ {
		nums[i], nums[current] = nums[current], nums[i]
		backtrackPermuteUnieque(current+1, nums, doubles)
		nums[i], nums[current] = nums[current], nums[i]
	}
}
