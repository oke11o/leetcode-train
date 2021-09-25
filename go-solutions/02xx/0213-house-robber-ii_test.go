package _2xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_rob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "",
			nums: []int{1},
			want: 1,
		},
		{
			name: "",
			nums: []int{2, 3, 2, 2},
			want: 5,
		},
		{
			name: "",
			nums: []int{2, 3, 2, 2, 3},
			want: 6,
		},
		{
			name: "",
			nums: []int{2, 3, 2},
			want: 3,
		},
		{
			name: "",
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
		{
			name: "",
			nums: []int{1, 2, 3},
			want: 3,
		},
		{
			name: "",
			nums: []int{1, 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rob(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
