package _8xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_peakIndexInMountainArray(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "",
			arr:  []int{0, 1, 0},
			want: 1,
		},
		{
			name: "",
			arr:  []int{0, 2, 1, 0},
			want: 1,
		},
		{
			name: "",
			arr:  []int{0, 10, 5, 2},
			want: 1,
		},
		{
			name: "",
			arr:  []int{3, 4, 5, 1},
			want: 2,
		},
		{
			name: "",
			arr:  []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := peakIndexInMountainArray(tt.arr)
			require.Equal(t, tt.want, got)
		})
	}
}
