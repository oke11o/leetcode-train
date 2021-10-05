package _7xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_floodFill(t *testing.T) {
	tests := []struct {
		name        string
		image       [][]int
		sr          int
		sc          int
		newColor    int
		want        [][]int
		explanation string
	}{
		{
			name:        "",
			image:       [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
			sr:          1,
			sc:          1,
			newColor:    2,
			want:        [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
			explanation: `From the center of the image with position (sr, sc) = (1, 1) (i.e., the red pixel), all pixels connected by a path of the same color as the starting pixel (i.e., the blue pixels) are colored with the new color. Note the bottom corner is not colored 2, because it is not 4-directionally connected to the starting pixel.`,
		},
		{
			name:     "",
			image:    [][]int{{0, 0, 0}, {0, 0, 0}},
			sr:       0,
			sc:       0,
			newColor: 2,
			want:     [][]int{{2, 2, 2}, {2, 2, 2}},
		},
		{
			name:     "",
			image:    [][]int{{0, 0, 0}, {0, 1, 1}},
			sr:       1,
			sc:       1,
			newColor: 1,
			want:     [][]int{{0, 0, 0}, {0, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := floodFill_bfs(tt.image, tt.sr, tt.sc, tt.newColor)
			require.Equal(t, tt.want, got)
		})
	}
}
