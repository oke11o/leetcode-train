package _1xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_hammingWeight(t *testing.T) {
	tests := []struct {
		name string
		num  uint32
		want int
	}{
		{
			name: "",
			num:  11,
			want: 3,
		},
		{
			name: "",
			num:  128, //00000000000000000000000010000000
			want: 1,
		},
		{
			name: "",
			num:  4294967293, //11111111111111111111111111111101
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hammingWeight(tt.num)
			require.Equal(t, tt.want, got)
		})
	}
}
