package _2xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_numSquares(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		want        int
		explanation string
	}{
		{
			name:        "",
			n:           12,
			want:        3,
			explanation: "12 = 4 + 4 + 4",
		},
		{
			name:        "",
			n:           13,
			want:        2,
			explanation: "13 = 4 + 9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numSquares(tt.n)
			require.Equal(t, tt.want, got)
		})
	}
}
