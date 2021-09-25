package _9xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_maxSubarraySumCircular(t *testing.T) {
	tests := []struct {
		name        string
		nums        []int
		want        int
		explanation string
	}{
		{
			name: "",
			nums: []int{5, 6, 1, 4, 8, -8, 7, -5, 3},
			want: 29,
		},
		{
			name: "",
			nums: []int{-5, 3, 5},
			want: 8,
		},
		{
			name: "",
			nums: []int{-9, 14, 24, -14, 12, 18, -18, -10, -10, -23, -2, -23, 11, 12, 18, -9, 9, -29, 12, 4, -8, 15, 11, -12, -16, -9, 19, -12, 22, 16},
			want: 99,
		},
		{
			name: "",
			nums: []int{-1, 3, -3, 9, -6, 8, -5, -5, -6, 10},
			want: 20,
		},
		{
			name:        "",
			nums:        []int{1, -2, 3, -2},
			want:        3,
			explanation: "Subarray [3] has maximum sum 3",
		},
		{
			name:        "",
			nums:        []int{5, -3, 5},
			want:        10,
			explanation: "Subarray [5,5] has maximum sum 5 + 5 = 10",
		},
		{
			name:        "",
			nums:        []int{3, -1, 2, -1},
			want:        4,
			explanation: "Subarray [2,-1,3] has maximum sum 2 + (-1) + 3 = 4",
		},
		{
			name:        "",
			nums:        []int{3, -2, 2, -3},
			want:        3,
			explanation: "Subarray [3] and [3,-2,2] both have maximum sum 3",
		},
		{
			name:        "",
			nums:        []int{-2, -3, -1},
			want:        -1,
			explanation: "Subarray [-1] has maximum sum -1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSubarraySumCircular(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
