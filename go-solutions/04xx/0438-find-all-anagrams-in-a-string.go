package _4xx

/**
 * 438. Find All Anagrams in a String
 * https://leetcode.com/problems/find-all-anagrams-in-a-string/
 * Medium
 * [#hash_table, #string, #sliding_window]
 *
 */
func findAnagrams(input string, pattern string) []int {
	paLength := len(pattern)
	in := []rune(input)
	paDict := make(map[rune]int, 26)
	for _, s := range pattern {
		paDict[s]++
	}
	result := []int{}
	inDict := make(map[rune]int, 26)
	for i := 0; i < len(in); i++ {
		if _, exists := paDict[in[i]]; exists {
			inDict[in[i]]++
		}
		if i >= paLength {
			if _, exists := paDict[in[i-paLength]]; exists {
				inDict[in[i-paLength]]--
			}
		}
		if isMapEq(inDict, paDict) {
			result = append(result, i+1-paLength)
		}
	}
	return result
}

func isMapEq(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
