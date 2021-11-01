package easy

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_moveZerosToLeft(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{
			name: "",
			in:   []int{1, 10, 20, 0, 59, 63, 0, 88, 0, 0, 0},
			want: []int{0, 0, 0, 0, 0, 1, 10, 20, 59, 63, 88},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := moveZerosToLeft(tt.in)
			require.Equal(t, tt.want, got)
		})
	}
}
