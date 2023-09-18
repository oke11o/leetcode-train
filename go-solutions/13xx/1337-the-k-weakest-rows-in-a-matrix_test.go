package _3xx

import (
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

/*
1337. The K Weakest Rows in a Matrix
https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix
Easy
Topics
Companies
Hint
You are given an m x n binary matrix mat of 1's (representing soldiers) and 0's (representing civilians). The soldiers are positioned in front of the civilians. That is, all the 1's will appear to the left of all the 0's in each row.

A row i is weaker than a row j if one of the following is true:

The number of soldiers in row i is less than the number of soldiers in row j.
Both rows have the same number of soldiers and i < j.
Return the indices of the k weakest rows in the matrix ordered from weakest to strongest.
*/
func kWeakestRows(mat [][]int, k int) []int {
	weights := make([][2]int, len(mat))
	for i := 0; i < len(mat); i++ {
		weights[i][0] = i
		l := 0
		r := len(mat[i]) - 1
		if mat[i][r] == 1 {
			weights[i][1] = r + 1
			continue
		}
		for l+1 < r {
			idx := (l + r) / 2
			if mat[i][idx] == 1 {
				l = idx
			} else {
				r = idx
			}
		}
		weights[i][0] = i
		if mat[i][l] == 1 {
			weights[i][1] = l + 1
		} else {
			weights[i][1] = l
		}
	}
	//[[1,0],[0,0],[1,0]]
	sort.Slice(weights, func(i, j int) bool {
		if weights[i][1] != weights[j][1] {
			return weights[i][1] < weights[j][1]
		}
		return weights[i][0] < weights[j][0]
	})

	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = weights[i][0]
	}

	return result
}

func Test_shortestPath(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		k    int
		want []int
	}{
		{
			name: "",
			mat:  [][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}},
			k:    3,
			want: []int{2, 0, 3},
		},
		{
			name: "",
			mat:  [][]int{{1, 0, 0, 0}, {1, 1, 1, 1}, {1, 0, 0, 0}, {1, 0, 0, 0}},
			k:    2,
			want: []int{0, 2},
		},
		{
			name: "",
			mat:  [][]int{{1, 0}, {0, 0}, {1, 0}},
			k:    2,
			want: []int{1, 0},
		},
		{
			name: "",
			mat:  [][]int{{1, 1, 0}, {1, 1, 0}, {1, 1, 1}, {1, 1, 1}, {0, 0, 0}, {1, 1, 1}, {1, 0, 0}},
			k:    6,
			want: []int{4, 6, 0, 1, 2, 3},
		},
		{
			name: "",
			mat:  [][]int{{1, 1, 1, 1, 1}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 1, 1, 1, 1}},
			k:    3,
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kWeakestRows(tt.mat, tt.k)
			require.Equal(t, tt.want, got)
		})
	}
}
