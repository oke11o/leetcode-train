package _0xx

/**
 * https://leetcode.com/problems/valid-parentheses/
 * 0020. Valid Parentheses
 * Easy
 * String, Stack
 */
func isValid(s string) bool {
	pairs := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}
	stack := []rune{}
	for _, v := range s {
		if _, ok := pairs[v]; ok {
			stack = append(stack, v)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		last := stack[len(stack)-1]
		if pairs[last] != v {
			stack = append(stack, v)
			continue

		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
