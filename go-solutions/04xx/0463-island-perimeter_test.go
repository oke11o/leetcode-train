package _4xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_islandPerimeter(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "",
			grid: [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}},
			want: 16,
		},
		{
			name: "",
			grid: [][]int{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}},
			want: 16,
		},
		{
			name: "",
			grid: [][]int{{1, 0, 1, 1}, {0, 0, 1, 0}, {0, 1, 1, 1}, {1, 0, 0, 1}},
			want: 24,
		},
		{
			name: "",
			grid: [][]int{{1}},
			want: 4,
		},
		{
			name: "",
			grid: [][]int{{1, 0}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := islandPerimeter_dfs(tt.grid)
			require.Equal(t, tt.want, got)
		})
	}
}
