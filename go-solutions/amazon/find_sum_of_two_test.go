package amazon

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_findSumOfTwo(t *testing.T) {
	tests := []struct {
		name string
		list []int
		sum  int
		want bool
	}{
		{
			name: "",
			list: []int{5, 7, 1, 2, 8, 4, 3},
			sum:  10,
			want: true,
		},
		{
			name: "",
			list: []int{5, 7, 1, 2, 8, 4, 3},
			sum:  19,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findSumOfTwo(tt.list, tt.sum)
			require.Equal(t, tt.want, got)
		})
	}
}
