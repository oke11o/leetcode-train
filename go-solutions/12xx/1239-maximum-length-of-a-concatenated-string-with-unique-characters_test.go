package _2xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_maxLength(t *testing.T) {
	tests := []struct {
		name        string
		arr         []string
		explanation string
		want        int
	}{
		{
			name:        "",
			arr:         []string{"un", "iq", "ue"},
			explanation: `All possible concatenations are "","un","iq","ue","uniq" and "ique".`,
			want:        4,
		},
		{
			name: "",
			arr:  []string{"un", "iqe", "ueabc"},
			want: 5,
		},
		{
			name:        "",
			arr:         []string{"cha", "r", "act", "ers"},
			explanation: `Possible solutions are "chaers" and "acters".`,
			want:        6,
		},
		{
			name: "",
			arr:  []string{"abcdefghijklmnopqrstuvwxyz"},
			want: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxLength(tt.arr)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_maxLength_binWords(t *testing.T) {
	tests := []struct {
		name  string
		word  string
		want  int
		want2 bool
	}{
		{
			name:  "",
			word:  "abcd",
			want:  15,
			want2: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := maxLength_binWords(tt.word)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.want2, got2)
		})
	}
}

func Test_maxLength_countBits(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{
			name: "",
			num:  1,
			want: 1,
		},
		{
			name: "",
			num:  2,
			want: 1,
		},
		{
			name: "",
			num:  3,
			want: 2,
		},
		{
			name: "",
			num:  4,
			want: 1,
		},
		{
			name: "",
			num:  7,
			want: 3,
		},
		{
			name: "",
			num:  79,
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxLength_countBits(tt.num)
			require.Equal(t, tt.want, got)
		})
	}
}
