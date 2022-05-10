package _0xx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
https://leetcode.com/problems/letter-combinations-of-a-phone-number/
17. Letter Combinations of a Phone Number
Medium
*/
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	result := make([]string, 0)

	m := map[uint8][]rune{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	var backtrack func(digits string, pos uint8, res []rune)
	backtrack = func(digits string, pos uint8, res []rune) {
		if len(res) == len(digits) { // pos == len(digits
			result = append(result, string(res))
			return
		}
		newRes := make([]rune, len(res)+1)
		copy(newRes, res)
		for _, s := range m[digits[pos]] {
			newRes[len(newRes)-1] = s
			backtrack(digits, pos+1, newRes)
		}
	}

	backtrack(digits, 0, []rune{})

	return result
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_letterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{digits: "23"},
			want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name: "",
			args: args{digits: ""},
			want: []string{},
		},
		{
			name: "",
			args: args{digits: "2"},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, letterCombinations(tt.args.digits), "letterCombinations(%v)", tt.args.digits)
		})
	}
}
