package _7xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_movesToChessboard(t *testing.T) {
	tests := []struct {
		name  string
		board [][]int
		want  int
	}{
		{
			name:  "",
			board: [][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {1, 0, 0, 1}, {1, 0, 0, 1}},
			want:  2,
		},
		{
			name:  "",
			board: [][]int{{0, 1}, {1, 0}},
			want:  0,
		},
		{
			name:  "",
			board: [][]int{{1, 0}, {1, 0}},
			want:  -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := movesToChessboard(tt.board)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_bitSlice2int(t *testing.T) {

	tests := []struct {
		name string
		bits []int
		want int
	}{
		{
			name: "",
			bits: []int{0},
			want: 0,
		},
		{
			name: "",
			bits: []int{1},
			want: 1,
		},
		{
			name: "",
			bits: []int{1, 0},
			want: 2,
		},
		{
			name: "",
			bits: []int{1, 1},
			want: 3,
		},
		{
			name: "",
			bits: []int{1, 0, 0},
			want: 4,
		},
		{
			name: "",
			bits: []int{1, 0, 1},
			want: 5,
		},
		{
			name: "",
			bits: []int{1, 1, 0},
			want: 6,
		},
		{
			name: "",
			bits: []int{0, 0, 1, 1, 0},
			want: 6,
		},
		{
			name: "",
			bits: []int{0, 1, 1, 0},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bitSlice2int(tt.bits)
			require.Equal(t, tt.want, got)
		})
	}
}
