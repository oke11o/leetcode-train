package _0xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mergeSortedArray(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{
			name:  "",
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:  "",
			nums1: []int{1, 2, 3, 7, 9, 0, 0, 0, 0},
			m:     5,
			nums2: []int{2, 3, 5, 6},
			n:     4,
			want:  []int{1, 2, 2, 3, 3, 5, 6, 7, 9},
		},
		{
			name:  "",
			nums1: []int{1},
			m:     1,
			nums2: []int{},
			n:     0,
			want:  []int{1},
		},
		{
			name:  "",
			nums1: []int{0},
			m:     0,
			nums2: []int{1},
			n:     1,
			want:  []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mergeSortedArray(tt.nums1, tt.m, tt.nums2, tt.n)
			require.Equal(t, tt.want, tt.nums1)
		})
	}
}
