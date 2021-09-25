package _0xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			name: "",
			nums: []int{1},
			want: 1,
		},
		{
			name: "",
			nums: []int{5, 4, -1, 7, 8},
			want: 23,
		},
		{
			name: "",
			nums: []int{-2, -7},
			want: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSubArray(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
