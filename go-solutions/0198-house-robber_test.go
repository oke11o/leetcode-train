package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_rob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "",
			nums: []int{2, 1, 1, 2},
			want: 4,
		},
		{
			name: "",
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
		{
			name: "",
			nums: []int{2, 7, 9, 3, 1},
			want: 12,
		},
		{
			name: "",
			nums: []int{2, 5, 3, 2, 5, 7, 98, 3, 1, 3, 45, 4, 6, 5, 75, 7},
			want: 235,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rob(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
