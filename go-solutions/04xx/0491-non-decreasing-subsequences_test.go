package _4xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_findSubsequences(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "",
			nums: []int{4, 6, 7, 7},
			want: [][]int{
				{4, 6},
				{4, 6, 7},
				{4, 6, 7, 7},
				{4, 7},
				{4, 7, 7},
				{6, 7},
				{6, 7, 7},
				{7, 7},
			},
		},
		// {
		// 	name: "",
		// 	nums: []int{4, 4, 3, 2, 1},
		// 	want: [][]int{
		// 		{4, 4},
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findSubsequences(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
