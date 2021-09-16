package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "",
			nums: []int{2, 2, 1},
			want: 1,
		},
		{
			name: "",
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			name: "",
			nums: []int{1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumber(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
