package _0xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_canJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "",
			nums: []int{2, 3, 1, 1, 4},
			want: true,
		},
		{
			name: "",
			nums: []int{3, 2, 1, 0, 4},
			want: false,
		},
		{
			name: "",
			nums: []int{3, 3, 1, 0, 4},
			want: true,
		},
		{
			name: "",
			nums: []int{3, 3, 1, 0, 1, 0, 1},
			want: false,
		},
		{
			name: "",
			nums: []int{3, 3, 1, 0, 2, 0, 1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canJump(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
