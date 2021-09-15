package go_solutions

func maxNumberOfBalloons(text string) int {
	var min = func(nn ...int) int {
		ans := 1<<((32<<(^uint(0)>>63))-1) - 1
		for _, n := range nn {
			if ans > n {
				ans = n
			}
		}
		return ans
	}
	var cntB, cntA, cntL, cntO, cntN int
	for _, v := range text {
		switch v {
		case 'b':
			cntB++
		case 'a':
			cntA++
		case 'l':
			cntL++
		case 'o':
			cntO++
		case 'n':
			cntN++
		}
	}
	return min(cntB, cntA, cntL/2, cntO/2, cntN)
}

func maxNumberOfBalloons_2(text string) int {
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
