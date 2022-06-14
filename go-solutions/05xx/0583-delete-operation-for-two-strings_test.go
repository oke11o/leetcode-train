package _5xx

import "testing"

/**
https://leetcode.com/problems/delete-operation-for-two-strings/
583. Delete Operation for Two Strings
Medium
*/
func minDistance(word1 string, word2 string) int {
	memo := make([][]int, len(word1)+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, len(word2)+1)
	}

	var lcs func(s1 string, s2 string, m int, n int, memo [][]int) int
	lcs = func(s1 string, s2 string, m int, n int, memo [][]int) int {
		if m == 0 || n == 0 {
			return 0
		}
		if memo[m][n] > 0 {
			return memo[m][n]
		}
		if s1[m-1] == s2[n-1] {
			memo[m][n] = 1 + lcs(s1, s2, m-1, n-1, memo)
		} else {
			l1 := lcs(s1, s2, m, n-1, memo)
			l2 := lcs(s1, s2, m-1, n, memo)
			if l1 > l2 {
				memo[m][n] = l1
			} else {
				memo[m][n] = l2
			}
		}
		return memo[m][n]

	}

	return len(word1) + len(word2) - 2*lcs(word1, word2, len(word1), len(word2), memo)
}

func minDistance_notWorking(word1 string, word2 string) int {
	m1 := make(map[rune]int)
	for _, v := range word1 {
		m1[v]++
	}
	m2 := make(map[rune]int)
	for _, v := range word2 {
		m2[v]++
	}

	result := 0
	for v, c1 := range m1 {
		c2 := m2[v]
		if c1 == c2 {
			continue
		}
		if c1 > c2 {
			result += c1 - c2
		} else {
			result += c2 - c1
		}
	}
	for v, c := range m2 {
		if _, ok := m1[v]; !ok {
			result += c
		}
	}
	return result
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_minDistance(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  int
	}{
		{
			name:  "",
			word1: "sea",
			word2: "eat",
			want:  2,
		},
		{
			name:  "",
			word1: "sea",
			word2: "ate",
			want:  4,
		},
		{
			name:  "",
			word1: "leetcode",
			word2: "etco",
			want:  4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.word1, tt.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
