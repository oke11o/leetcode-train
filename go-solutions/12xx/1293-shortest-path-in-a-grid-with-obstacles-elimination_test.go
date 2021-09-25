package _2xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_shortestPath(t *testing.T) {
	tests := []struct {
		name        string
		grid        [][]int
		k           int
		want        int
		explanation string
	}{
		{
			name: "",
			grid: [][]int{{0, 0, 0},
				{1, 1, 0},
				{0, 0, 0},
				{0, 1, 1},
				{0, 0, 0}},
			k:    1,
			want: 6,
			explanation: `
The shortest path without eliminating any obstacle is 10.
The shortest path with one obstacle elimination at position (3,2) is 6. Such path is (0,0) -> (0,1) -> (0,2) -> (1,2) -> (2,2) -> (3,2) -> (4,2).`,
		},
		{
			name: "",
			grid: [][]int{{0, 1, 1},
				{1, 1, 1},
				{1, 0, 0}},
			k:    1,
			want: -1,
			explanation: `
We need to eliminate at least two obstacles to find such a walk.
`,
		},
		{
			name: "",
			grid: [][]int{{0, 1, 1},
				{1, 1, 1},
				{1, 0, 0}},
			k:    2,
			want: 4,
			explanation: `
We need to eliminate at least two obstacles to find such a walk.
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shortestPath(tt.grid, tt.k)
			require.Equal(t, tt.want, got)
		})
	}
}
