package _0xx

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
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			name:   "",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 7,
			want:   3,
		},
		{
			name:   "",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 6,
			want:   2,
		},
		{
			name:   "",
			nums:   []int{0, 1, 2, 4, 5, 6, 7},
			target: 0,
			want:   0,
		},
		{
			name:   "",
			nums:   []int{0, 1, 2, 4, 5, 6, 7},
			target: 7,
			want:   6,
		},
		{
			name:   "",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		{
			name:   "",
			nums:   []int{0},
			target: 0,
			want:   0,
		},
		{
			name:   "",
			nums:   []int{},
			target: 1,
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := search(tt.nums, tt.target)
			require.Equal(t, tt.want, got)
		})
	}
}
