package _1xx

/**
https://leetcode.com/problems/longest-substring-with-at-most-two-distinct-characters/
159. Longest Substring with At Most Two Distinct Characters
Medium
#sliding_window
*/
/**
corner tc:
- empty string
- uniq letters
- only one letter
*/
func lengthOfLongestSubstringTwoDistinct(s string) int {
	ans := 0
	first := 0
	second := 0
	m := make(map[byte]int)
	length := 0 // ece ba
	for first < len(s) {
		firstSym := s[first] // a
		length++             // 3
		m[firstSym]++        // {e:1, b:1, a:1}
		for len(m) > 2 {     //3>2  //for condition when we should move second pointer
			secondSym := s[second] // e
			m[secondSym]--         // {e:1, c:0, b:1}
			length--               //2
			if m[secondSym] == 0 { //0==0
				delete(m, secondSym) // {e:1, b:1}
			}
			second++ // ->2
		}
		if len(m) <= 2 && ans < length { //2==2 3<2 true
			ans = length // 3
		}
		first++
	}
	return ans // 3
}
