package sum_of_subsets

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_sumOfSubsets(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		sum  int
		want [][]int
	}{
		{
			name: "",
			in:   []int{5, 10, 12, 13, 15, 18},
			sum:  100,
			want: nil,
		},
		{
			name: "",
			in:   []int{5, 10, 12, 13, 15, 18},
			sum:  30,
			want: [][]int{{5, 10, 15}, {5, 12, 13}, {12, 18}},
		},
		{
			name: "",
			in:   []int{2, 3, 6, 7},
			sum:  7,
			want: [][]int{{7}}, //{2,2,3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumOfSubsets(tt.in, tt.sum)
			require.Equal(t, tt.want, got)
		})
	}
}
