package backtracking

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_queensProblems(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]int
	}{
		{
			name: "",
			n:    4,
			want: [][]int{{2, 4, 1, 3}, {3, 1, 4, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := queensProblems(tt.n)
			require.Equal(t, tt.want, got)
		})
	}
}
