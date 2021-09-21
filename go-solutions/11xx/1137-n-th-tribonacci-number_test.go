package _1xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_tribonacci(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "",
			n:    4,
			want: 4,
		},
		{
			name: "",
			n:    25,
			want: 1389537,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tribonacci(tt.n)
			require.Equal(t, tt.want, got)
		})
	}
}
