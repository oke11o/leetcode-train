package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_totalFruit(t *testing.T) {
	tests := []struct {
		name   string
		fruits []int
		want   int
	}{
		{
			name:   "",
			fruits: []int{1, 2, 1},
			want:   3,
			//Explanation: We can pick from all 3 trees.
		},

		{
			name:   "",
			fruits: []int{0, 1, 2, 2},
			want:   3,
			//Explanation: We can pick from trees [1,2,2].
		},

		{
			name:   "",
			fruits: []int{1, 2, 3, 2, 2},
			want:   4,
			//Explanation: We can pick from trees [2,3,2,2].
		},

		{
			name:   "",
			fruits: []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4},
			want:   5,
			//Explanation: We can pick from trees [1,2,1,1,2].
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := totalFruit(tt.fruits)
			require.Equal(t, tt.want, got)
		})
	}
}
