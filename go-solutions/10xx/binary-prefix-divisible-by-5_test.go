package _0xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_prefixesDivBy5(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []bool
	}{
		{
			name: "",
			nums: []int{0, 1, 1, 1, 1, 1},
			want: []bool{true, false, false, false, true, false},
		},
		{
			name: "",
			nums: []int{0, 1, 1},
			want: []bool{true, false, false},
		},
		{
			name: "",
			nums: []int{1, 1, 1},
			want: []bool{false, false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := prefixesDivBy5(tt.nums)
			require.Equal(t, tt.want, got)

		})
	}
}
