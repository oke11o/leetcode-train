package _0xx

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := [][]int{intervals[0]}
	for i := 0; i < len(intervals); i++ {
		val := intervals[i]
		lastIdx := len(result) - 1
		if result[lastIdx][1] >= val[0] {
			if result[lastIdx][1] < val[1] {
				result[lastIdx][1] = val[1]
			}
		} else {
			result = append(result, val)
		}
	}
	return result
}

func Test_merge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{
			name:      "",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:      [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "",
			intervals: [][]int{{1, 4}, {4, 5}},
			want:      [][]int{{1, 5}},
		},
		{
			name:      "",
			intervals: [][]int{{1, 4}, {0, 4}},
			want:      [][]int{{0, 4}},
		},
		{
			name:      "",
			intervals: [][]int{{1, 4}, {2, 3}},
			want:      [][]int{{1, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := merge(tt.intervals)
			require.Equal(t, tt.want, got)
		})
	}
}
