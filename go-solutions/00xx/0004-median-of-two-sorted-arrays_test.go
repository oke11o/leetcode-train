package _0xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			name:  "",
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2.0,
		},
		{
			name:  "",
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2.0,
		},
		{
			name:  "",
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMedianSortedArrays(tt.nums1, tt.nums2)
			require.Equalf(t, tt.want, result, "findMedianSortedArrays(%v, %v)", tt.nums1, tt.nums2)
		})
	}
}
