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
	left := 0
	result := 0

	district := make(map[uint8]int)
	sum := 0
	for right := 0; right < len(input); right++ {
		sum++
		rightLetter := input[right]
		district[rightLetter]++
		for sum > len(district) {
			leftLetter := input[left]
			district[leftLetter]--
			if district[leftLetter] == 0 {
				delete(district, leftLetter)
			}
			sum--
			left++
		}

		if result < sum {
			result = sum
		}
	}

	return result
}
