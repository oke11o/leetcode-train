package _8xx

import "testing"

/**
809. Expressive Words
Medium
https://leetcode.com/problems/expressive-words/
[#array, #two_pointers, #string]
*/
func expressiveWords(s string, words []string) int {
	cnt := 0
	for _, w := range words {
		if expressiveWord(s, w) {
			cnt++
		}
	}
	return cnt
}

func expressiveWord(str string, word string) bool {
	strP := 0
	wordP := 0
	for strP < len(str) && wordP < len(word) {
		if str[strP] != word[wordP] {
			return false
		}
		strLen := 1
		for strP < len(str)-1 && str[strP] == str[strP+1] {
			strLen++
			strP++
		}
		wordLen := 1
		for wordP < len(word)-1 && word[wordP] == word[wordP+1] {
			wordLen++
			wordP++
		}
		if (strLen < 3 && wordLen != strLen) || (strLen >= 3 && strLen < wordLen) {
			return false
		}
		wordP++
		strP++

	}
	return strP == len(str) && wordP == len(word)
}

func Test_expressiveWords(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		words []string
		want  int
	}{
		{
			name:  "",
			s:     "heeellooo",
			words: []string{"hello", "hi", "helo"},
			want:  1,
		},
		{
			name:  "",
			s:     "zzzzzyyyyy",
			words: []string{"zzyy", "zy", "zyy"},
			want:  3,
		},
		{
			name:  "",
			s:     "abcd",
			words: []string{"abc"},
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expressiveWords(tt.s, tt.words); got != tt.want {
				t.Errorf("expressiveWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expressiveWord(t *testing.T) {
	tests := []struct {
		name string
		s    string
		word string
		want bool
	}{
		{
			name: "",
			s:    "heeellooo",
			word: "hello",
			want: true,
		},
		{
			name: "",
			s:    "heeellooo",
			word: "hi",
			want: false,
		},
		{
			name: "",
			s:    "heeellooo",
			word: "helo",
			want: false,
		},
		{
			name: "",
			s:    "zzzzzyyyyy",
			word: "zzyy",
			want: true,
		},
		{
			name: "",
			s:    "zzzzzyyyyy",
			word: "zy",
			want: true,
		},
		{
			name: "",
			s:    "zzzzzyyyyy",
			word: "zyy",
			want: true,
		},
		{
			name: "",
			s:    "abcd",
			word: "abc",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expressiveWord(tt.s, tt.word); got != tt.want {
				t.Errorf("expressiveWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
