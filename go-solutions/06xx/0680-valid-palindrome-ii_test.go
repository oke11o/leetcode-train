package _6xx

import "testing"

/**
680. Valid Palindrome II
Easy
https://leetcode.com/problems/valid-palindrome-ii/
*/
func validPalindrome(s string) bool {
	var isPali = func(s string) bool {
		left := 0
		right := len(s) - 1
		for left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
		return true
	}
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
			continue
		}
		return isPali(s[left:right]) || isPali(s[left+1:right+1])
	}
	return true
}

func Test_validPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "",
			s:    "aba",
			want: true,
		},
		{
			name: "",
			s:    "abca",
			want: true,
		},
		{
			name: "",
			s:    "abc",
			want: false,
		},
		{
			name: "",
			s:    "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPalindrome(tt.s); got != tt.want {
				t.Errorf("validPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
