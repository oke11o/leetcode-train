package _5xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

/*
1503. Last Moment Before All Ants Fall Out of a Plank
https://leetcode.com/problems/last-moment-before-all-ants-fall-out-of-a-plank/
*/
func getLastMoment(n int, left []int, right []int) int {
	minF := func(in []int) int {
		min := 999999999
		for _, v := range in {
			if min > v {
				min = v
			}
		}
		return min
	}
	maxF := func(in []int) int {
		max := -1
		for _, v := range in {
			if max < v {
				max = v
			}
		}
		return max
	}

	leftDist := maxF(left)
	rightMin := minF(right)
	rightDist := n - rightMin
	if len(left) == 0 {
		return rightDist
	}
	if len(right) == 0 {
		return leftDist
	}
	if leftDist > rightDist {
		return leftDist
	}
	return rightDist
}

func Test_getLastMoment(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		left  []int
		right []int
		want  int
	}{
		{
			name:  "",
			n:     4,
			left:  []int{4, 3},
			right: []int{0, 1},
			want:  4,
		},
		{
			name:  "",
			n:     7,
			left:  []int{},
			right: []int{0, 1, 2, 3, 4, 5, 6, 7},
			want:  7,
		},
		{
			name:  "",
			n:     7,
			left:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			right: []int{},
			want:  7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getLastMoment(tt.n, tt.left, tt.right)
			require.Equal(t, tt.want, got)
		})
	}
}
