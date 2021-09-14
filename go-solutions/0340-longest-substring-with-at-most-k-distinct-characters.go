package go_solutions

func lengthOfLongestSubstringKDistinct(input string, k int) int {
	result := 0
	left := 0

	distinct := make(map[uint8]int)
	for right := 0; right < len(input); right++ {
		rightCh := input[right]
		distinct[rightCh]++

		for len(distinct) > k {
			leftCh := input[left]
			distinct[leftCh]--
			if distinct[leftCh] == 0 {
				delete(distinct, leftCh)
			}
			left++
		}
		if result < right-left+1 {
			result = right - left + 1
		}
	}

	return result
}
