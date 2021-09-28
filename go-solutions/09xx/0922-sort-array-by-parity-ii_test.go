package _9xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_sortArrayByParityII(t *testing.T) {
	tests := []struct {
		name        string
		nums        []int
		want        []int
		explanation string
	}{
		{
			name: "",
			nums: []int{0, 1, 2, 1, 4, 1, 3, 4},
			want: []int{0, 1, 2, 1, 4, 1, 4, 3},
		},
		{
			name: "",
			nums: []int{4, 1, 1, 0, 1, 0},
			want: []int{4, 1, 0, 1, 0, 1},
		},
		{
			name: "",
			nums: []int{888, 505, 627, 846},
			want: []int{888, 505, 846, 627},
		},
		{
			name:        "",
			nums:        []int{4, 2, 5, 7},
			want:        []int{4, 5, 2, 7},
			explanation: "[4,7,2,5], [2,5,4,7], [2,7,4,5] would also have been accepted.",
		},
		{
			name:        "",
			nums:        []int{2, 3},
			want:        []int{2, 3},
			explanation: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortArrayByParityII(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
