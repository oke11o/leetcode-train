package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

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
