package _2xx

func maxLength(arr []string) int {
	m := make(map[int32][]int)
	for i, word := range arr {
		for _, symbol := range word {
			m[symbol] = append(m[symbol], i)
		}
	}

	return 0
}
