package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_lengthOfLongestSubstringKDistinct(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			name: "",
			s:    "eceba",
			k:    3,
			want: 4,
		},
		{
			name: "",
			s:    "eceba",
			k:    2,
			want: 3,
		},
		{
			name: "",
			s:    "WORLD",
			k:    4,
			want: 4,
		},
		{
			name: "",
			s:    "assddfffdasdffdfasfasdfadffdfwertnbv",
			k:    4,
			want: 29,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstringKDistinct(tt.s, tt.k)
			require.Equal(t, tt.want, got)
		})
	}
}
