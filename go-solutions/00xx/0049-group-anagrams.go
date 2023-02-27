package _0xx

/*
49. Group Anagrams
Medium
*/
func groupAnagrams(strs []string) [][]string {
	hash := func(in string) [26]int {
		result := [26]int{}
		for _, b := range in {
			result[b-'a']++
		}

		return result
	}

	m := make(map[[26]int][]string)
	for _, str := range strs {
		h := hash(str)
		m[h] = append(m[h], str)
	}

	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}

	return result
}
