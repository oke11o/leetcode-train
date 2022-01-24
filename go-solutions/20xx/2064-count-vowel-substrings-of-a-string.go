package _0xx

import "strconv"

/**
 * 2062. Count Vowel Substrings of a String
 * https://leetcode.com/problems/count-vowel-substrings-of-a-string/
 * Easy
 * Tags: #hash_table, #string, #sliding_window
 *
 * Cool solutions:
 * - https://leetcode.com/problems/frequency-of-the-most-frequent-element/discuss/1175088/C%2B%2B-Maximum-Sliding-Window-Cheatsheet-Template!
 * - https://github.com/lzl124631x/LeetCode/tree/master/leetcode/239.%20Sliding%20Window%20Maximum
 */
func countVowelSubstrings(in string) int {
	word := []rune(in)
	result, startVowel, left, right, distinctVowels := 0, 0, 0, 0, 0
	vowels := map[rune]int{'a': 0, 'e': 0, 'i': 0, 'o': 0, 'u': 0}

	for ; right < len(word); right++ {
		rightChar := word[right]
		rCharT := strconv.QuoteRune(rightChar)
		_ = rCharT
		if _, ok := vowels[rightChar]; ok {
			vowels[rightChar]++
			if vowels[rightChar] == 1 {
				distinctVowels++
			}
			for distinctVowels == 5 {
				leftChar := word[left]
				vowels[leftChar]--
				if vowels[leftChar] == 0 {
					distinctVowels--
				}
				left++
			}
			result = result + (left - startVowel)
		} else {
			vowels = map[rune]int{'a': 0, 'e': 0, 'i': 0, 'o': 0, 'u': 0}
			distinctVowels = 0
			left, startVowel = right+1, right+1
		}
	}

	return result
}
