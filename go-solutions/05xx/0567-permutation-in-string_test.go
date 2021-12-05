package _5xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_checkInclusion(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{
			name: "",
			s1:   "eidbaooo",
			s2:   "ab",
			want: false,
		},
		{
			name: "",
			s1:   "ab",
			s2:   "eidbaooo",
			want: true,
		},
		{
			name: "",
			s1:   "ab",
			s2:   "eidboaoo",
			want: false,
		},
		{
			name: "",
			s1:   "adc",
			s2:   "dcda",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkInclusion(tt.s1, tt.s2)
			require.Equal(t, tt.want, got)
		})
	}
}
