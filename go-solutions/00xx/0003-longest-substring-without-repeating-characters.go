package _0xx

// 0003. Longest Substring Without Repeating Characters
// Medium
// Similar 0340. Longest Substring with At Most K Distinct Characters*
// Hash Table, String, Sliding Window
// https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(input string) int {
	if len(input) == 0 {
		return 0
	}
	right := 0
	result := 0

	district := make(map[uint8]int)
	sum := 0
	for left := 0; left < len(input); left++ {
		sum++
		leftSym := input[left]
		if _, ok := district[leftSym]; !ok {
			district[leftSym] = 0
		}
		district[leftSym]++
		for sum > len(district) {
			rightSym := input[right]
			district[rightSym]--
			if district[rightSym] == 0 {
				delete(district, rightSym)
			}
			sum--
			right++
		}

		if result < sum {
			result = sum
		}
	}

	return result
}
