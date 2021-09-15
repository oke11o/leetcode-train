package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_sumSquareDigits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "",
			n:    19,
			want: 82,
		},
		{
			name: "",
			n:    824,
			want: 84,
		},
		{
			name: "",
			n:    2,
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumSquareDigits(tt.n)
			require.Equal(t, tt.want, got)
		})
	}
}
func Test_isHappy(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{
			name: "",
			n:    19,
			want: true,
		},
		{
			name: "",
			n:    2,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isHappy(tt.n)
			require.Equal(t, tt.want, got)
		})
	}
}
