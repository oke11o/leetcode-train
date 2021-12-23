package _5xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_change(t *testing.T) {
	tests := []struct {
		name   string
		amount int
		coins  []int
		want   int
	}{
		{
			name:   "",
			amount: 5,
			coins:  []int{1, 2, 5},
			want:   4,
		},
		{
			name:   "",
			amount: 7,
			coins:  []int{1, 2, 5},
			want:   6,
		},
		{
			name:   "",
			amount: 3,
			coins:  []int{2},
			want:   0,
		},
		{
			name:   "",
			amount: 10,
			coins:  []int{10},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := change(tt.amount, tt.coins)
			require.Equal(t, tt.want, got)
		})
	}
}
