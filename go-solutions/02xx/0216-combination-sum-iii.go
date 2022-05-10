package _2xx

/**
https://leetcode.com/problems/combination-sum-iii/
216. Combination Sum III
Medium
*/
func combinationSum3(k int, n int) [][]int {
	result := [][]int{}

	var backtrack func(nextDigit int, curDigits []int, curSum int)
	backtrack = func(nextDigit int, curDigits []int, curSum int) {
		if curSum == n {
			if len(curDigits) == k {
				result = append(result, curDigits)
			}
			return
		}
		if len(curDigits) >= k {
			return
		}
		dgts := make([]int, len(curDigits)+1)
		copy(dgts, curDigits)
		for i := nextDigit + 1; i < 10; i++ {
			newSum := curSum + i
			if newSum > n {
				break
			}
			dgts[len(dgts)-1] = i
			backtrack(i, dgts, newSum)
		}
	}
	backtrack(0, []int{}, 0)

	return result
}
