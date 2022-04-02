package _6xx

import "testing"

/**
680. Valid Palindrome II
Easy
https://leetcode.com/problems/valid-palindrome-ii/

В начале решал с передачей в helper function isPali()  строку.
Но в решениях предложено передвать индексы и строку. Но так как мы можем сделать функцию 2го порядка. Можно использовать замыкание.
Код стал чуть проще
*/
func validPalindrome(s string) bool {
	var checkPalindrome = func(i, j int) bool {
		for i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}

	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return checkPalindrome(i+1, j) || checkPalindrome(i, j-1)
		}
		i++
		j--
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
