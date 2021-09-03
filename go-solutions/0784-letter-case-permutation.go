package go_solutions

func letterCasePermutation(s string) []string {
	result := []string{s}
	for i := 0; i < len(s); i++ {
		if isSymbol(s[i]) {
			length := len(result)
			for j := 0; j < length; j++ {
				word := result[j]
				result = append(result, word[:i]+string(alterSymbol(s[i]))+word[i+1:])
			}
		}
	}

	return result
}

func isSymbol(u uint8) bool {
	if u >= 97 && u <= 122 {
		return true
	}
	if u >= 65 && u <= 90 {
		return true
	}
	return false
}

func alterSymbol(u uint8) uint8 {
	if u >= 97 && u <= 122 {
		return u - 32
	}
	return u + 32
}
