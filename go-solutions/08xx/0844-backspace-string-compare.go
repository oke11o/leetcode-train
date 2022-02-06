package _8xx

import "strings"

/**
 *0844. Backspace String Compare
 *Easy
 ***Tags:** [Two Pointers, String, Stack, Simulation]
 * https://leetcode.com/problems/backspace-string-compare/
 */
func backspaceCompare(s string, t string) bool {
	var backspace = func(s string, r uint8) string {
		var result []string
		cntToDel := 0
		for i := len(s) - 1; i >= 0; i-- {
			v := s[i]
			if v == r {
				cntToDel++
			} else {
				if cntToDel == 0 {
					result = append([]string{string(v)}, result...)
				} else {
					cntToDel--
				}
			}
		}
		return strings.Join(result, "")
	}

	return backspace(s, '#') == backspace(t, '#')
}

func backspaceCompare2(s string, t string) bool {
	return backspace(s, '#') == backspace(t, '#')
}

func backspace(s string, back rune) string {
	ss := []rune(s)
	result := make([]rune, len(ss))
	cnt := 0
	curCnt := 0
	rIdx := len(ss) - 1
	for i := len(ss) - 1; i >= 0; i-- {
		if ss[i] == back {
			curCnt++
			cnt++
		} else if curCnt > 0 {
			curCnt--
			cnt++
		} else {
			result[rIdx] = ss[i]
			rIdx--
		}
	}
	return string(result[cnt:])
}
