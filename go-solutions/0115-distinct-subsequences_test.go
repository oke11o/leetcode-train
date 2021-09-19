package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_numDistinct(t *testing.T) {
	tests := []struct {
		name   string
		str    string
		target string
		want   int
	}{
		{
			name:   "",
			str:    "rabbbit",
			target: "rabbit",
			want:   3,
		},
		{
			name:   "",
			str:    "babgbag",
			target: "bag",
			want:   5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numDistinct(tt.str, tt.target)
			require.Equal(t, tt.want, got)
		})
	}
}
