package _2xx

/**
https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string-ii/
1209. Remove All Adjacent Duplicates in String II
Medium
*/
func removeDuplicates(s string, k int) string {
	if len(s) == 0 {
		return s
	}
	// "deeedbbcccbdaa"
	//            ^
	result := [][]byte{{}} // {d}

	for i := 0; i < len(s); i++ {
		v := s[i] // b
		lastIdx := len(result) - 1
		if len(result[lastIdx]) != 0 && v == result[lastIdx][0] {
			result[lastIdx] = append(result[lastIdx], v) // {dd}{bb}
			if len(result[lastIdx]) == k {
				result = result[0:lastIdx] // {d}
			}
		} else {
			result = append(result, []byte{v}) // {dd}{b}
		}
	}

	res := []byte{}
	for _, r := range result {
		res = append(res, r...)
	}

	return string(res)
}
