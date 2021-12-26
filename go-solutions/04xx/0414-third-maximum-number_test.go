package _4xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_thirdMax(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "",
			nums: []int{3, 2, 1},
			want: 1,
		},
		{
			name: "",
			nums: []int{1, 2},
			want: 2,
		},
		{
			name: "",
			nums: []int{2, 2, 3, 1},
			want: 1,
		},
		{
			name: "",
			nums: []int{1, 2, 2, 5, 3, 5},
			want: 2,
		},
		{
			name: "",
			nums: []int{1, 2, -2147483648},
			want: -2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := thirdMax(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
