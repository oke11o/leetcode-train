package go_solutions

func reverseOnlyLetters(input string) string {
	left := 0
	right := len(input) - 1
	result := make([]byte, len(input))
	for left < right {
		leftCh := input[left]
		rightCh := input[right]
		if !isSymbol(leftCh) {
			result[left] = leftCh
			left++
			continue
		}
		if !isSymbol(rightCh) {
			result[right] = rightCh
			right--
			continue
		}
		result[left] = rightCh
		result[right] = leftCh
		left++
		right--
	}
	if left == right {
		result[left] = input[left]
	}
	return string(result)
}

/*
func isSymbol(u uint8) bool {
	if u >= 97 && u <= 122 {
		return true
	}
	if u >= 65 && u <= 90 {
		return true
	}
	return false
}
*/
