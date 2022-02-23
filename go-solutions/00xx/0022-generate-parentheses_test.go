package _0xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var generateParenthesis_result []string

/**
 * 0022. Generate Parentheses
 * Medium
 * #amazon
 * String, Dynamic Programming, Backtracking
 * https://leetcode.com/problems/generate-parentheses/
 * https://www.educative.io/m/all-possible-braces
 * https://www.educative.io/blog/crack-amazon-coding-interview-questions
 */
func generateParenthesis(n int) []string {
	generateParenthesis_result = []string{}

	generateParenthesis_rec(n, 0, 0, "")

	return generateParenthesis_result
}

func generateParenthesis_rec(n, opened, closed int, current string) {
	if opened+closed == 2*n {
		generateParenthesis_result = append(generateParenthesis_result, current)
		return
	}
	if opened < n {
		generateParenthesis_rec(n, opened+1, closed, current+"(")
	}
	if closed < opened {
		generateParenthesis_rec(n, opened, closed+1, current+")")
	}
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_generateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			name: "",
			n:    1,
			want: []string{"()"},
		},
		{
			name: "",
			n:    2,
			want: []string{"(())", "()()"},
		},
		{
			name: "",
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "",
			n:    4,
			want: []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"},
		},
		{
			name: "",
			n:    5,
			want: []string{"((((()))))", "(((()())))", "(((())()))", "(((()))())", "(((())))()", "((()(())))", "((()()()))", "((()())())", "((()()))()", "((())(()))", "((())()())", "((())())()", "((()))(())", "((()))()()", "(()((())))", "(()(()()))", "(()(())())", "(()(()))()", "(()()(()))", "(()()()())", "(()()())()", "(()())(())", "(()())()()", "(())((()))", "(())(()())", "(())(())()", "(())()(())", "(())()()()", "()(((())))", "()((()()))", "()((())())", "()((()))()", "()(()(()))", "()(()()())", "()(()())()", "()(())(())", "()(())()()", "()()((()))", "()()(()())", "()()(())()", "()()()(())", "()()()()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateParenthesis(tt.n)
			require.Equal(t, tt.want, got)
		})
	}
}
