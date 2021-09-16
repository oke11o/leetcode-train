package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNumArray_SumRange(t *testing.T) {
	tests := []struct {
		name  string
		sum   []int
		left  int
		right int
		want  int
	}{
		{
			name:  "",
			sum:   []int{-2, 0, 3, -5, 2, -1},
			left:  0,
			right: 2,
			want:  1,
		},
		{
			name:  "",
			sum:   []int{-2, 0, 3, -5, 2, -1},
			left:  2,
			right: 5,
			want:  -1,
		},
		{
			name:  "",
			sum:   []int{-2, 0, 3, -5, 2, -1},
			left:  0,
			right: 5,
			want:  -3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Constructor(tt.sum)
			got := a.SumRange(tt.left, tt.right)
			require.Equal(t, tt.want, got)
		})
	}
}
