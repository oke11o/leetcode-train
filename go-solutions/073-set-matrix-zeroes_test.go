package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_setZeroes(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name:   "",
			matrix: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			want:   [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			name:   "",
			matrix: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			want:   [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.NotEqual(t, tt.want, tt.matrix)
			setZeroes(tt.matrix)
			require.Equal(t, tt.want, tt.matrix)
		})
	}
}
