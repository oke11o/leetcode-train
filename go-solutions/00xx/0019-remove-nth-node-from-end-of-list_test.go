package _0xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeNthFromEnd(t *testing.T) {
	tests := []struct {
		name string
		head []int
		n    int
		want []int
	}{
		{
			name: "",
			head: []int{1, 2, 3, 4, 5},
			n:    2,
			want: []int{1, 2, 3, 5},
		},
		{
			name: "",
			head: []int{1},
			n:    1,
			want: []int{},
		},
		{
			name: "",
			head: []int{1, 2},
			n:    1,
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := BuildList(tt.head)
			got := removeNthFromEnd(head, tt.n)
			require.Equal(t, BuildList(tt.want), got)
		})
	}
}
