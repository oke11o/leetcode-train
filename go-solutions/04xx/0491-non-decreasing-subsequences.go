package _4xx

/*
491. Non-decreasing Subsequences
Medium
2.4K
179
company
Amazon
company
Yahoo
Given an integer array nums, return all the different possible non-decreasing subsequences of the given array with at least two elements.
You may return the answer in any order.
*/
func findSubsequences(nums []int) [][]int {
	result := [][]int{}
	var backtrack func(start int, sequence []int)
	backtrack = func(start int, oldSeq []int) {
		sequence := make([]int, len(oldSeq))
		copy(sequence, oldSeq)
		if len(sequence) == 0 || sequence[len(sequence)-1] <= nums[start] {
			sequence = append(sequence, nums[start])
			if len(sequence) > 1 {
				result = append(result, sequence)
			}
		}
		for i := start + 1; i < len(nums); i++ {
			backtrack(i, sequence)
		}

	}

	backtrack(0, []int{})

	return result
}
