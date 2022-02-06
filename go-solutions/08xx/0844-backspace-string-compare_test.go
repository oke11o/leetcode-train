package _8xx

import (
	"testing"

	"github.com/stretchr/testify/require"
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
			got := backspaceCompare2(tt.s, tt.t)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_backspace(t *testing.T) {
	tests := []struct {
		name string
		s    string
		back rune
		want string
	}{
		{
			name: "",
			s:    "a##c",
			back: '#',
			want: "c",
		},
		{
			name: "",
			s:    "#a#c",
			back: '#',
			want: "c",
		},
		{
			name: "",
			s:    "ab##",
			back: '#',
			want: "",
		},
		{
			name: "",
			s:    "ab#######",
			back: '#',
			want: "",
		},
		{
			name: "",
			s:    "asdffwefew##awf3#faf#####",
			back: '#',
			want: "asdffwefa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspace(tt.s, tt.back); got != tt.want {
				t.Errorf("backspace() = %v, want %v", got, tt.want)
			}
		})
	}
}
