package _0xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	prev := 1
	result := 1
	for i := 2; i <= n; i++ {
		t := result
		result = t + prev
		prev = t
	}
	return result
}

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
