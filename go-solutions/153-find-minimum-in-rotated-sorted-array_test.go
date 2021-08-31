package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_findMinInRotatedSortedArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "",
			nums: []int{3, 4, 5, 1, 2},
			want: 1,
		},
		{
			name: "",
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			want: 0,
		},
		{
			name: "",
			nums: []int{11, 13, 15, 17},
			want: 11,
		},
		{
			name: "",
			nums: []int{44},
			want: 44,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMinInRotatedSortedArray(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
