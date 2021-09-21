package _5xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_fib(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "",
			n:    2,
			want: 1,
		},
		{
			name: "",
			n:    3,
			want: 2,
		},
		{
			name: "",
			n:    4,
			want: 3,
		},
		{
			name: "",
			n:    5,
			want: 5,
		},
		{
			name: "",
			n:    6,
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fib(tt.n)
			require.Equal(t, tt.want, got)
		})
	}
}
