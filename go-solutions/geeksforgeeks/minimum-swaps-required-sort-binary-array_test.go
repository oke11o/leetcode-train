package geeksforgeeks

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMinSwaps(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "",
			input: []int{0, 1, 0, 1},
			want:  1,
		},
		{
			name:  "",
			input: []int{0, 0, 1, 0, 1, 0, 1, 1},
			want:  3,
		},
		{
			name:  "",
			input: []int{0, 1, 0, 1, 0},
			want:  3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//got := findMinSwaps_bruteForce(tt.input)
			got := findMinSwaps(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
