package _0xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_rotate(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name:   "",
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			name:   "",
			matrix: [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			want:   [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
		{
			name:   "",
			matrix: [][]int{{1}},
			want:   [][]int{{1}},
		},
		{
			name:   "",
			matrix: [][]int{{1, 2}, {3, 4}},
			want:   [][]int{{3, 1}, {4, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.matrix)
			require.Equal(t, tt.want, tt.matrix)
		})
	}
}

func Test_transpose(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name:   "",
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transpose(tt.matrix)
			require.Equal(t, tt.want, tt.matrix)
		})
	}
}

func Test_reverse(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name:   "",
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   [][]int{{3, 2, 1}, {6, 5, 4}, {9, 8, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reflect(tt.matrix)
			require.Equal(t, tt.want, tt.matrix)
		})
	}
}
