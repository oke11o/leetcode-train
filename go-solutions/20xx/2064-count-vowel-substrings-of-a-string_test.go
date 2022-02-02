package _0xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countVowelSubstrings(t *testing.T) {
	tests := []struct {
		name string
		word string
		want int
	}{
		{
			name: "",
			word: "xxaiioueaiixx",
			want: 12,
		},
		{
			name: "",
			word: "xxaiioueiiaxx",
			want: 7,
		},
		{
			name: "",
			word: "aeiouu",
			want: 2,
		},
		{
			name: "",
			word: "unicornarihan",
			want: 0,
		},
		{
			name: "",
			word: "cuaieuouac",
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countVowelSubstrings(tt.word)
			require.Equal(t, tt.want, got)
		})
	}
}
