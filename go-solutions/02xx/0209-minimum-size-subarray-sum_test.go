package _2xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minSubArrayLen(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		want   int
	}{
		{
			name:   "",
			target: 11,
			nums:   []int{1, 2, 3, 4, 5},
			want:   3,
		},
		{
			name:   "",
			target: 7,
			nums:   []int{2, 3, 1, 2, 4, 3},
			want:   2,
		},
		{
			name:   "",
			target: 4,
			nums:   []int{1, 4, 4},
			want:   1,
		},
		{
			name:   "",
			target: 11,
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minSubArrayLen(tt.target, tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
