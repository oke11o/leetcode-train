package _4xx

/**
 * 438. Find All Anagrams in a String
 * https://leetcode.com/problems/find-all-anagrams-in-a-string/
 * Medium
 * [#hash_table, #string, #sliding_window]
 *
 * Ознакомиться с алгоритмом поиска подстроки Rabin-Karp https://en.wikipedia.org/wiki/Rabin%E2%80%93Karp_algorithm
 */
func findAnagrams_two_maps(input string, pattern string) []int {
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

func findAnagrams_wrong(input string, pattern string) []int {
	count := len(pattern)
	in := []rune(input)
	dict := make(map[rune]int, 26)
	for _, s := range pattern {
		dict[s]++
	}
	result := []int{}
	p1 := 0
	for p2 := 0; p2 < len(in); p2++ {
		if _, exists := dict[in[p2]]; exists {
			dict[in[p2]]--
			count--
		}
		if p2-p1 == len(pattern) {
			if _, exists := dict[in[p1]]; exists {
				dict[in[p1]]++
				count++
			}
			p1++
		}
		if count == 0 {
			result = append(result, p1)
		}
	}
	return result
}
