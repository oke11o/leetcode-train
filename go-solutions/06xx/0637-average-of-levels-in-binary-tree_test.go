package _6xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_averageOfLevels(t *testing.T) {
	tests := []struct {
		name string
		node []int
		want []float64
	}{
		{
			name: "",
			node: []int{3, 9, 20, nilTreeNodeVal, nilTreeNodeVal, 15, 7},
			want: []float64{3.00000, 14.50000, 11.00000},
		},
		{
			name: "",
			node: []int{3, 9, 20, 15, 7},
			want: []float64{3.00000, 14.50000, 11.00000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := averageOfLevels(createTreeNodeFromSlice(tt.node))
			require.Equal(t, tt.want, got)
		})
	}
}
