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
func longestPalindrome_bruteForce(s string) string {
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

/**
Тут просто.
Первым циклом бежим по всем символам.
На каждом шаге пытаемся найти максимально длинный палиндром, считая, что i - это центр палиндрома
ВАЖНО: центр может быть посередине (aba, abba). Поэтому еще проверяем втору позицию.
TC: O(n^2)
*/
func longestPalindrome(s string) string {
	var expandAroundCenter = func(left, right int) int {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		return right - left - 1
	}

	if len(s) < 2 {
		return s
	}
	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(i, i)
		len2 := expandAroundCenter(i, i+1)
		length := len1
		if length < len2 {
			length = len2
		}
		if length > end-start {
			start = i - (length-1)/2
			end = i + length/2
		}
	}
	return s[start : end+1]
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
			//want: "bab",
			want: "aba",
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
