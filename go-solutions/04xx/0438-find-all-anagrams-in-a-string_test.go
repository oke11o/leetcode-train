package _4xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findAnagrams(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want []int
	}{
		{
			name: "",
			s:    "cbaebabacd",
			p:    "abc",
			want: []int{0, 6},
		},
		{
			name: "",
			s:    "abab",
			p:    "ab",
			want: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findAnagrams(tt.s, tt.p)
			require.Equal(t, tt.want, got)
		})
	}
}
