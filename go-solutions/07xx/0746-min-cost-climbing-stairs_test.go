package _7xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_minCostClimbingStairs(t *testing.T) {
	tests := []struct {
		name string
		cost []int
		want int
	}{
		{
			name: "",
			cost: []int{10, 15, 20},
			want: 15,
		},
		{
			name: "",
			cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			want: 6,
		},
		{
			name: "",
			cost: []int{1, 2, 3, 2},
			want: 4,
		},
		{
			name: "",
			cost: []int{1, 2, 3, 2, 1},
			want: 4,
		},
		{
			name: "",
			cost: []int{1, 2},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minCostClimbingStairs(tt.cost)
			require.Equal(t, tt.want, got)
		})
	}
}
