package _7xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_nextGreatestLetter(t *testing.T) {
	tests := []struct {
		name    string
		letters []byte
		target  byte
		want    byte
	}{
		{
			name:    "",
			letters: []byte{'c', 'f', 'j'},
			target:  'a',
			want:    'c',
		},
		{
			name:    "",
			letters: []byte{'c', 'f', 'j'},
			target:  'c',
			want:    'f',
		},
		{
			name:    "",
			letters: []byte{'c', 'f', 'j'},
			target:  'd',
			want:    'f',
		},
		{
			name:    "",
			letters: []byte{'c', 'f', 'j'},
			target:  'g',
			want:    'j',
		},
		{
			name:    "",
			letters: []byte{'c', 'f', 'j'},
			target:  'j',
			want:    'c',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nextGreatestLetter(tt.letters, tt.target)
			require.Equal(t, tt.want, got)
		})
	}
}
