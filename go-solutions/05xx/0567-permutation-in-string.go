package _5xx

// 0567. Permutation in String
// Medium
// Hash Table, Two Pointers, String, Sliding Window
// https://leetcode.com/problems/permutation-in-string/
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	symbols1 := make([]int, 26)
	symbols2 := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		v1 := s1[i] - 'a'
		symbols1[v1]++
		v2 := s2[i] - 'a'
		symbols2[v2]++
	}
	for i := 0; i < len(s2)-len(s1); i++ {
		if checkInclusion_match(symbols1, symbols2) {
			return true
		}
		j := len(s1) + i
		iLeft := s2[i] - 'a'
		iRight := s2[j] - 'a'
		symbols2[iLeft]--
		symbols2[iRight]++
	}

	return checkInclusion_match(symbols1, symbols2)
}

func checkInclusion_match(symbols1 []int, symbols2 []int) bool {
	for i, v1 := range symbols1 {
		if v1 != symbols2[i] {
			return false
		}
	}
	return true
}
