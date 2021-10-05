package _0xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_jump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "",
			nums: []int{2, 3, 1, 1, 4},
			want: 2,
		},
		{
			name: "",
			nums: []int{2, 3, 0, 1, 4},
			want: 2,
		},
		{
			name: "",
			nums: []int{2, 3, 1, 1, 2, 1, 1},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := jump(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
