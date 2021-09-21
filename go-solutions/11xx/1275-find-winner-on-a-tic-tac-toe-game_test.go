package _1xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_tictactoe(t *testing.T) {
	tests := []struct {
		name  string
		moves [][]int
		want  string
	}{
		{
			name:  "",
			moves: [][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}},
			want:  "A",
		},
		{
			name:  "",
			moves: [][]int{{0, 0}, {1, 1}, {0, 1}, {0, 2}, {1, 0}, {2, 0}},
			want:  "B",
		},
		{
			name:  "",
			moves: [][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}},
			want:  "Draw",
		},
		{
			name:  "",
			moves: [][]int{{0, 0}, {1, 1}},
			want:  "Pending",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tictactoe(tt.moves)
			require.Equal(t, tt.want, got)
		})
	}
}
