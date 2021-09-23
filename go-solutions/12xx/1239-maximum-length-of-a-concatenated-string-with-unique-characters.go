package _2xx

func maxLength(arr []string) int {

	result := 0
	stack := []int{0}
	resLen := len(stack)
	for i := 0; i < len(arr); i++ {
		binWord, dubExists := maxLength_binWords(arr[i])
		if dubExists {
			continue
		}
		for i := 0; i < resLen; i++ {
			inRes := stack[i]
			isUniq := inRes & binWord
			if isUniq != 0 {
				continue
			}
			newWord := inRes | binWord
			stack = append(stack, newWord)
			resLen = len(stack)
			binCnt := maxLength_countBits(newWord)
			if result < binCnt {
				result = binCnt
			}

		}
	}

	return result
}

func maxLength_binWords(word string) (int, bool) {
	result := 0
	hasDubs := false
	for _, sym := range word {
		sym = sym - 'a'
		t := 1 << sym
		dub := result & t
		if dub > 0 {
			hasDubs = true
		}
		result |= t
	}
	return result, hasDubs
}

func maxLength_countBits(n int) int {
	result := 0
	for n != 0 {
		if n%2 == 1 {
			result++
		}
		n = n / 2
	}
	return result
}
