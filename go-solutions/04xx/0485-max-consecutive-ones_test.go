package _4xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_findMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "",
			nums: []int{1, 1, 0, 1, 1, 1},
			want: 3,
		},
		{
			name: "",
			nums: []int{1, 0, 1, 1, 0, 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxConsecutiveOnes(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
