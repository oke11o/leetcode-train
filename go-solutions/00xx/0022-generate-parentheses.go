package _0xx

var generateParenthesis_result []string

/**
 * 0022. Generate Parentheses
 * Medium
 * String, Dynamic Programming, Backtracking
 * https://leetcode.com/problems/generate-parentheses/
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
