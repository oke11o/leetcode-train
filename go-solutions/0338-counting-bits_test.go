package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_countBits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "",
			n:    2,
			want: []int{0, 1, 1},
		},
		{
			name: "",
			n:    5,
			want: []int{0, 1, 1, 2, 1, 2},
		},
		{
			name: "",
			n:    7,
			want: []int{0, 1, 1, 2, 1, 2, 2, 3},
		},
		{
			name: "",
			n:    8,
			want: []int{0, 1, 1, 2, 1, 2, 2, 3, 1},
		},
		{
			name: "",
			n:    9,
			want: []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countBits(tt.n)
			require.Equal(t, tt.want, got)
		})
	}
}
