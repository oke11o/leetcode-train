package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_findDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "",
			nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			want: []int{2, 3},
		},
		{
			name: "",
			nums: []int{1, 1, 2},
			want: []int{1},
		},
		{
			name: "",
			nums: []int{1},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDuplicates(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
