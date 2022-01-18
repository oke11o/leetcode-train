package _2xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_smallestDivisor(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		threshold int
		want      int
	}{
		{
			name:      "",
			nums:      []int{1, 2, 5, 9},
			threshold: 6,
			want:      5,
		},
		{
			name:      "",
			nums:      []int{44, 22, 33, 11, 1},
			threshold: 5,
			want:      44,
		},
		{
			name:      "",
			nums:      []int{4},
			threshold: 5,
			want:      1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := smallestDivisor(tt.nums, tt.threshold)
			require.Equal(t, tt.want, got)

		})
	}
}
