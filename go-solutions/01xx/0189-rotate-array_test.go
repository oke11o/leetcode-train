package _1xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_rotate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "",
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    3,
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "",
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    9,
			want: []int{6, 7, 1, 2, 3, 4, 5},
		},
		{
			name: "",
			nums: []int{-1, -100, 3, 99},
			k:    2,
			want: []int{3, 99, -1, -100},
		},
		{
			name: "",
			nums: []int{1},
			k:    0,
			want: []int{1},
		},
		{
			name: "",
			nums: []int{1},
			k:    2,
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.nums, tt.k)
			require.Equal(t, tt.want, tt.nums)
		})
	}
}
