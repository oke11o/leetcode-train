package _0xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_maxScoreSightseeingPair(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{
			name:   "",
			values: []int{8, 1, 5, 2, 6},
			want:   11,
		},
		{
			name:   "",
			values: []int{1, 2},
			want:   2,
		},
		{
			name:   "",
			values: []int{1, 8, 2, 1, 1, 1, 9, 5, 4},
			want:   13,
		},
		{
			name:   "",
			values: []int{1, 10, 2, 1, 1, 1, 9, 5, 4},
			want:   14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxScoreSightseeingPair(tt.values)
			require.Equal(t, tt.want, got)
		})
	}
}
