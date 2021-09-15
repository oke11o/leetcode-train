package go_solutions

func maxNumberOfBalloons(text string) int {
	assert := "balloon"
	set := make(map[rune]int)
	for _, v := range assert {
		set[v]++
	}
	dict := make(map[rune]int)
	for _, v := range text {
		if _, ok := set[v]; ok {
			dict[v]++
		}
	}
	if len(dict) < len(set) {
		return 0
	}

	for s, count := range set {
		dict[s] = dict[s] / count
	}

	result := 1<<((32<<(^uint(0)>>63))-1) - 1
	for _, s := range dict {
		if result > s {
			result = s
		}
	}

	return result
}
