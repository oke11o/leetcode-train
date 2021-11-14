package graph_coloring

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_graphColoring(t *testing.T) {
	tests := []struct {
		name string
		in   Graph
		want []string
	}{
		{
			name: "",
			in:   Graph{0: {1, 2, 3}, 1: {0, 3, 4}, 2: {0, 3, 4}, 3: {0, 1, 2}, 4: {1, 2}},
			want: []string{
				"RGGBR", "RGGBB", "RBBGR", "RBBGG",
				"GRRBG", "GRRBB", "GBBRR", "GBBRG",
				"BRRGG", "BRRGB", "BGGRR", "BGGRB",
			},
		},
		{
			name: "",
			in:   Graph{0: {1, 2}, 1: {0, 3}, 2: {0, 3}, 3: {1, 2}},
			want: []string{
				"RGGR", "RGGB", "RGBR", "RBGR", "RBBR", "RBBG",
				"GRRG", "GRRB", "GRBG", "GBRG", "GBBR", "GBBG",
				"BRRG", "BRRB", "BRGB", "BGRB", "BGGR", "BGGB",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := graphColoring(tt.in)
			require.ElementsMatch(t, tt.want, got)
		})
	}
}

func Test_boundFunc(t *testing.T) {
	tests := []struct {
		name     string
		in       Graph
		idx      int
		curColor rune
		curState []rune
		want     bool
	}{
		{
			name:     "",
			in:       Graph{0: {1, 2}, 1: {0, 3}, 2: {0, 3}, 3: {1, 2}},
			curState: []rune{0, 0, 0, 0},
			idx:      0,
			curColor: 42,
			want:     true,
		},
		{
			name:     "",
			in:       Graph{0: {1, 2}, 1: {0, 3}, 2: {0, 3}, 3: {1, 2}},
			curState: []rune{42, 0, 0, 0},
			idx:      1,
			curColor: 42,
			want:     false,
		},
		{
			name:     "",
			in:       Graph{0: {1, 2}, 1: {0, 3}, 2: {0, 3}, 3: {1, 2}},
			curState: []rune{42, 0, 0, 0},
			idx:      3,
			curColor: 42,
			want:     true,
		},
		{
			name:     "",
			in:       Graph{0: {1, 2}, 1: {0, 3}, 2: {0, 3}, 3: {1, 2}},
			curState: []rune{0, 42, 0, 0},
			idx:      3,
			curColor: 42,
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := boundFunc(tt.in, tt.idx, tt.curColor, tt.curState)
			require.Equal(t, tt.want, got)
		})
	}
}
