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

func checkInclusion_Optimized(s1 string, s2 string) bool {
	const alfabetCnt = 26
	if len(s1) > len(s2) {
		return false
	}

	symCounts1 := make([]int, alfabetCnt)
	symCounts2 := make([]int, alfabetCnt)
	for i := 0; i < len(s1); i++ {
		v1 := s1[i] - 'a'
		symCounts1[v1]++
		v2 := s2[i] - 'a'
		symCounts2[v2]++
	}
	equalCnt := 0
	for i := 0; i < alfabetCnt; i++ {
		if symCounts1[i] == symCounts2[i] {
			equalCnt++
		}
	}
	for i := 0; i < len(s2)-len(s1); i++ {
		if equalCnt == alfabetCnt {
			return true
		}

		j := len(s1) + i
		iLeft := s2[i] - 'a'
		iRight := s2[j] - 'a'

		symCounts2[iRight]++
		if symCounts1[iRight] == symCounts2[iRight] { // Правый символ дошел до нужного числа
			equalCnt++
		} else if symCounts1[iRight] == symCounts2[iRight]-1 { // А до этого прибавления, он был равен?
			equalCnt--
		}

		symCounts2[iLeft]--
		if symCounts1[iLeft] == symCounts2[iLeft] { // Левый символ дошел до нужного числа
			equalCnt++
		} else if symCounts1[iLeft] == symCounts2[iLeft]+1 { // А до этого вычитания, он был равен?
			equalCnt--
		}
	}

	return equalCnt == alfabetCnt
}

func checkInclusion_match(symbols1 []int, symbols2 []int) bool {
	for i, v1 := range symbols1 {
		if v1 != symbols2[i] {
			return false
		}
	}
	return true
}
