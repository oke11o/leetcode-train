package _8xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_backspaceCompare(t *testing.T) {
	tests := []struct {
		name        string
		s           string
		t           string
		explanation string
		want        bool
	}{
		{
			name:        "",
			s:           "ab#c",
			t:           "ad#c",
			explanation: `Both s and t become "ac"`,
			want:        true,
		},
		{
			name:        "",
			s:           "ab##",
			t:           "c#d#",
			explanation: `Both s and t become ""`,
			want:        true,
		},
		{
			name:        "",
			s:           "a##c",
			t:           "#a#c",
			explanation: `Both s and t become "c"`,
			want:        true,
		},
		{
			name:        "",
			s:           "a#c",
			t:           "b",
			explanation: `s becomes "c" while t becomes "b"`,
			want:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := backspaceCompare(tt.s, tt.t)
			require.Equal(t, tt.want, got)
		})
	}
}
