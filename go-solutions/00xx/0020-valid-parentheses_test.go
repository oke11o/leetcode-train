package _0xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_isValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "",
			s:    "()",
			want: true,
		},
		{
			name: "",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "",
			s:    "(]",
			want: false,
		},
		{
			name: "",
			s:    "([)]",
			want: false,
		},
		{
			name: "",
			s:    "((([]{[]}){}[{}]))",
			want: true,
		},
		{
			name: "",
			s:    "{[]}",
			want: true,
		},
		{
			name: "",
			s:    "]",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.s)
			require.Equal(t, tt.want, got)
		})
	}
}
