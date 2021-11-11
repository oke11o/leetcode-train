package _0xx

import "sort"

/**
 * 0039. Combination Sum
 * Medium
 * Array, Backtracking
 * https://leetcode.com/problems/combination-sum/
 */
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result = make([][]int, 0)
	backtrack(candidates, 0, target, 0, []int{})
	return result
}

var result [][]int

func backtrack(in []int, from int, needSum, prevSum int, current []int) {
	for i := from; i < len(in); i++ {
		val := in[i]
		newSum := prevSum + val
		if newSum == needSum {
			tmp := make([]int, len(current))
			copy(tmp, current)
			tmp = append(tmp, val)
			result = append(result, tmp)
			return
		}
		if newSum < needSum {
			tmp := make([]int, len(current))
			copy(tmp, current)
			tmp = append(tmp, val)
			backtrack(in, i, needSum, newSum, tmp)
		}
	}
}
