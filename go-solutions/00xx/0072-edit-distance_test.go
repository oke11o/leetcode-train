package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
https://leetcode.com/problems/edit-distance/
hard
72. Edit Distance
#string, #dynamic_programming
*/

/*
-	- 	h   o   r   s   e
-   0   1   2   3   4   5
r   1   1   2   2   3   4
o   2   2   1   2   3   4
s   3   3   0   0   0   0

dp[3][2]
hor cast to ro - one replace and one deletion
h cat to ros - two deletions and one replace

-	-   i   n   t   e   n   t   i   o   n
-	0   1   2   3   4   5   6   7   8   9
e   1   1   2   3   3   4   5   6   7   8
x   2   2   2   3   4   4   5   6   7   8
e   3   3   3   3   3   4   5   6   7   8
c   4   4   4   4   4   4   5   6   7   8
u   5   5   5   5   5   5   5   6   7   8
t   6   6   6   5   6   6   5   6   7   8
i   7   6   7   8   9   10  11  5   6   7
o   8   7   8   9   10  11  12  13  5   6
n   9   8   7   8   9   10  11  12  13  5
*/
func minDistance(word1 string, word2 string) int { // horse, ros
	min := func(a, b, c int) int {
		if a > b {
			a = b
		}
		if a > c {
			return c
		}
		return a
	}

	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] { //is same characters
				//no operation
				dp[i][j] = dp[i-1][j-1]
			} else {
				// one of operations
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
				//replace    delete      insert
			}
		}
	}

	return dp[m][n]
}

func Test_minDistance(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  int
	}{
		{
			name:  "",
			word1: "horse",
			word2: "ros",
			want:  3,
		},
		{
			name:  "",
			word1: "intention",
			word2: "execution",
			want:  5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minDistance(tt.word1, tt.word2), "minDistance(%v, %v)", tt.word1, tt.word2)
		})
	}
}
