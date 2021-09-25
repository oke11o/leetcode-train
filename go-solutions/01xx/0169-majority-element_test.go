package _1xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// 0169. Majority Element
func Test_majorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "",
			nums: []int{3, 2, 3},
			want: 3,
		},
		{
			name: "",
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := majorityElement(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
