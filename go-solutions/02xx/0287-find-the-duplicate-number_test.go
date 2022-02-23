package _2xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "",
			nums: []int{1, 3, 4, 2, 2},
			want: 2,
		},
		{
			name: "",
			nums: []int{3, 1, 3, 4, 2},
			want: 3,
		},
		{
			name: "",
			nums: []int{1, 1},
			want: 1,
		},
		{
			name: "",
			nums: []int{1, 1, 2},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDuplicate(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
