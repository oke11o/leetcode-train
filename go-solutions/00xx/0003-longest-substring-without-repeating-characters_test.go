package _0xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "",
			s:    "bbbbb",
			want: 1,
		},
		{
			name: "",
			s:    "pwwkew",
			want: 3,
		},
		{
			name: "",
			s:    "",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.s)
			require.Equal(t, tt.want, got)
		})
	}
}
