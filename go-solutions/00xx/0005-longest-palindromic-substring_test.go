package _0xx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
https://leetcode.com/problems/longest-palindromic-substring/
5. Longest Palindromic Substring
Medium

TC: O(n^3)
*/
func longestPalindrome(s string) string {
	isPalindrome := func(i, j int) bool {
		for i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}
	left := -1
	right := -1
	maxLength := 0
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j >= i; j-- {
			if isPalindrome(i, j) {
				length := j - i + 1 // NOTICE: important for 1 letter strings
				if maxLength < length {
					maxLength = length
					left = i
					right = j
				}
			}
		}
	}
	if left == -1 {
		return ""
	}
	return s[left : right+1]
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "",
			s:    "babab",
			want: "babab",
		},
		{
			name: "",
			s:    "babad",
			want: "bab",
		},
		{
			name: "",
			s:    "cbbd",
			want: "bb",
		},
		{
			name: "",
			s:    "a",
			want: "a",
		},
		{
			name: "",
			s:    "aa",
			want: "aa",
		},
		{
			name: "",
			s:    "aaa",
			want: "aaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, longestPalindrome(tt.s), "longestPalindrome(%v)", tt.s)
		})
	}
}
