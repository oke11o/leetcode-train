package _1xx

import "testing"

/**
125. Valid Palindrome
Easy
https://leetcode.com/problems/valid-palindrome/
*/
func isPalindrome(s string) bool {
	var convert = func(v uint8) (uint8, bool) {
		if v >= 'A' && v <= 'Z' {
			return v + 'a' - 'A', true
		}
		if (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
			return v, true
		}

		return v, false
	}

	left := 0
	right := len(s) - 1
	for left < right {
		leftSym, leftOk := convert(s[left])
		rightSym, rightOk := convert(s[right])
		if !leftOk {
			left++
			continue
		}
		if !rightOk {
			right--
			continue
		}
		if leftSym != rightSym {
			return false
		}
		left++
		right--
	}
	return true
}

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "amanaplanacanalpanama",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "raceacar",
			s:    "race a car",
			want: false,
		},
		{
			name: " ",
			s:    " ",
			want: true,
		},
		{
			name: " ",
			s:    "0p",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
