package _2xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_maxLength(t *testing.T) {
	type args struct {
	}
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
			want: 4,
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
