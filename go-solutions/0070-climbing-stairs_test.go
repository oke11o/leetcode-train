package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_climbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "",
			n:    1,
			want: 1,
		},
		{
			name: "",
			n:    2,
			want: 2,
		},
		{
			name: "",
			n:    3,
			want: 3,
		},
		{
			name: "",
			n:    4,
			want: 5,
		},
		{
			name: "",
			n:    5,
			want: 8,
		},
		{
			name: "",
			n:    6,
			want: 13,
		},
		{
			name: "",
			n:    7,
			want: 21,
		},
		{
			name: "",
			n:    8,
			want: 34,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbStairs(tt.n)
			require.Equal(t, tt.want, got)
		})
	}
}
