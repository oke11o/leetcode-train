package yandex

func findAnag(in, need []rune) int {
	if len(in) < len(need) {
		return -1
	}
	if len(need) == 0 {
		return 0
	}
	M := make(map[rune]int)
	// init map
	for _, v := range need {
		M[v]++
	} // {o:1,w:1,r:1}
	// Algorithm
	right := 0
	left := 0
	// left < 14 - 3 + 1 = 12; left = [0, 11]
	for left < len(in)-len(need) {
		for right-left <= len(need) {
			M[in[right]]--
			if M[in[right]] == 0 {
				delete(M, in[right])
			}
			if len(M) == 0 {
				return left
			}
			right++
		}
		M[in[left]]++
		if M[in[left]] == 0 {
			delete(M, in[left])
		}
		left++
		if len(M) == 0 {
			return left
		}
	}

	return -1
}
