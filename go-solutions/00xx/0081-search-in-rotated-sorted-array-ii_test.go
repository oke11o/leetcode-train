package _0xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_search81(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		{
			name:   "",
			nums:   []int{0, 1, 1, 2, 0, 0},
			target: 2,
			want:   true,
		},
		{
			name:   "",
			nums:   []int{1, 3},
			target: 0,
			want:   false,
		},
		{
			name:   "",
			nums:   []int{1, 1},
			target: 2,
			want:   false,
		},
		{
			name:   "",
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			target: 2,
			want:   false,
		},
		{
			name:   "",
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 13, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			target: 13,
			want:   true,
		},
		{
			name:   "",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   true,
		},
		{
			name:   "",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 7,
			want:   true,
		},
		{
			name:   "",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 6,
			want:   true,
		},
		{
			name:   "",
			nums:   []int{0, 1, 2, 4, 5, 6, 7},
			target: 0,
			want:   true,
		},
		{
			name:   "",
			nums:   []int{0, 1, 2, 4, 5, 6, 7},
			target: 7,
			want:   true,
		},
		{
			name:   "",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   false,
		},
		{
			name:   "",
			nums:   []int{0},
			target: 0,
			want:   true,
		},
		{
			name:   "",
			nums:   []int{},
			target: 1,
			want:   false,
		},
		{
			name:   "",
			nums:   []int{1, 1, 1, 1, 2, 1, 1, 1, 1},
			target: 2,
			want:   true,
		},
		{
			name:   "",
			nums:   []int{1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1},
			target: 2,
			want:   true,
		},
		{
			name:   "",
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1},
			target: 2,
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := search81(tt.nums, tt.target)
			require.Equal(t, tt.want, got)
		})
	}
}
