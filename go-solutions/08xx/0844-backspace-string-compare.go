package _8xx

import "strings"

// 0844. Backspace String Compare
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
