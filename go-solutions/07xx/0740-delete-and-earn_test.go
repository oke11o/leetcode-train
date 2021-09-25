package _7xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_deleteAndEarn(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "",
			nums: []int{3, 1},
			want: 4,
		},
		{
			name: "",
			nums: []int{3, 4, 2},
			want: 6,
		},
		{
			name: "",
			nums: []int{2, 2, 3, 3, 3, 4},
			want: 9,
		},
		{
			name: "",
			nums: []int{2, 2, 3, 3, 3, 4, 5, 5, 5, 5, 6, 6, 6, 7, 7, 7, 7, 7, 7, 7, 7, 8, 8, 8, 9, 9, 9, 9, 9},
			want: 130,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteAndEarn(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
