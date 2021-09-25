package _2xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_majorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "",
			nums: []int{1, 1, 2, 2, 7, 7, 8, 8, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3},
			want: []int{3, 9},
		},
		{
			name: "",
			nums: []int{3, 3},
			want: []int{3},
		},
		{
			name: "",
			nums: []int{0, 0, 0},
			want: []int{0},
		},
		{
			name: "",
			nums: []int{3, 2, 3},
			want: []int{3},
		},
		{
			name: "",
			nums: []int{1, 2},
			want: []int{2, 1},
		},
		{
			name: "",
			nums: []int{1},
			want: []int{1},
		},
		{
			name: "",
			nums: []int{1, 1, 1, 3, 3, 3, 2, 2, 2},
			want: []int{},
		},
		{
			name: "",
			nums: []int{1, 1, 1, 3, 3, 2, 2, 2},
			want: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := majorityElement(tt.nums)
			require.ElementsMatch(t, tt.want, got)
		})
	}
}
