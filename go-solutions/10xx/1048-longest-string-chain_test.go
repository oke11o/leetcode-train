package _0xx

import "testing"

/**
https://leetcode.com/problems/longest-string-chain/
1048. Longest String Chain
Medium
*/
func longestStrChain(words []string) int {
	wordSet := make(map[string]struct{}, len(words))
	for _, w := range words {
		wordSet[w] = struct{}{}
	}
	memo := make(map[string]int, len(words))
	var dfs func(word string) int
	dfs = func(word string) int {
		if m, ok := memo[word]; ok {
			return m
		}

		// This stores the maximum length of word sequence possible with the 'currentWord' as the
		maxLength := 1

		// creating all possible strings taking out one character at a time from the `currentWord`
		currentWord := []rune(word)
		for i := 0; i < len(currentWord); i++ {
			newRunes := make([]rune, len(word)-1)
			copy(newRunes, currentWord[0:i])
			if i < len(currentWord)-1 {
				copy(newRunes[i:], currentWord[i+1:])
			}
			newWord := string(newRunes)
			if _, exists := wordSet[newWord]; exists {
				currentLength := 1 + dfs(newWord)
				if maxLength < currentLength {
					maxLength = currentLength
				}
			}
		}
		memo[word] = maxLength
		return maxLength
	}
	ans := 0
	for _, w := range words {
		r := dfs(w)
		if ans < r {
			ans = r
		}
	}

	return ans
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_longestStrChain(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  int
	}{
		{
			name:  "",
			words: []string{"a", "b", "ba", "bca", "bda", "bdca"},
			want:  4,
		},
		{
			name:  "",
			words: []string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"},
			want:  5,
		},
		{
			name:  "",
			words: []string{"abcd", "dbqca"},
			want:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestStrChain(tt.words); got != tt.want {
				t.Errorf("longestStrChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
