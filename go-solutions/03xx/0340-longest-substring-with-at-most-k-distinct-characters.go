package _3xx

/**
https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/
340. Longest Substring with At Most K Distinct Characters
Medium
#sliding_window
*/
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	ans := 0
	first := 0
	second := 0

	//   _
	// eceba
	//    ^
	m := make(map[byte]int)
	length := 0
	for first < len(s) {
		sym1 := s[first]
		m[sym1]++
		length++

		for len(m) > k {
			sym2 := s[second]
			m[sym2]--
			if m[sym2] == 0 {
				delete(m, sym2)
			}
			length--
			second++
		}

		if len(m) <= k && ans < length {
			ans = length
		}
		first++
	}

	return ans
}
