package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_search(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			want:   4,
		},
		{
			name:   "",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			want:   -1,
		},
		{
			name:   "",
			nums:   []int{5},
			target: 5,
			want:   0,
		},
		{
			name:   "",
			nums:   []int{4, 5},
			target: 5,
			want:   1,
		},
		{
			name:   "",
			nums:   []int{4, 5},
			target: 4,
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := search(tt.nums, tt.target)
			require.Equal(t, tt.want, got)
		})
	}
}
