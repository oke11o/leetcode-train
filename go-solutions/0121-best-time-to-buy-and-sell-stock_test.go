package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			name:   "",
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.prices)
			require.Equal(t, tt.want, got)
		})
	}
}
