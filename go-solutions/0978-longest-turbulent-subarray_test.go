package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_maxTurbulenceSize(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "",
			arr:  []int{9, 9},
			want: 1,
		},
		{
			name: "",
			arr:  []int{9, 4, 2, 10, 7, 8, 8, 1, 9},
			want: 5,
		},
		{
			name: "",
			arr:  []int{9, 4, 2, 10, 7, 8, 1, 9},
			want: 7,
		},
		{
			name: "",
			arr:  []int{2, 3, 7, 5, 6, 3, 9, 5, 6, 4, 4, 4, 34, 6, 56, 4, 67, 54, 345, 54, 345, 54, 54},
			want: 11,
		},
		{
			name: "",
			arr:  []int{4, 8, 12, 16},
			want: 2,
		},
		{
			name: "",
			arr:  []int{100},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxTurbulenceSize(tt.arr)
			require.Equal(t, tt.want, got)
		})
	}
}
