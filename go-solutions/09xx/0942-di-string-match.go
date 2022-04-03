package _9xx

/**
https://leetcode.com/problems/di-string-match/
942. DI String Match
Easy
#greedy
*/
func diStringMatch(s string) []int {
	result := make([]int, 0, len(s)+1)
	lo := 0
	hi := len(s)
	for _, l := range s {
		if l == 'I' {
			result = append(result, lo)
			lo++
		} else {
			result = append(result, hi)
			hi--
		}
	}
	result = append(result, lo)

	return result
}
