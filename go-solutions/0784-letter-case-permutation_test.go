package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_letterCasePermutation(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{
			name: "",
			s:    "C",
			want: []string{"C", "c"},
		},
		{
			name: "",
			s:    "aZ",
			want: []string{"aZ", "AZ", "az", "Az"},
		},
		{
			name: "",
			s:    "az",
			want: []string{"az", "Az", "aZ", "AZ"},
		},
		{
			name: "",
			s:    "a1b2",
			want: []string{"a1b2", "A1b2", "a1B2", "A1B2"},
		},
		{
			name: "",
			s:    "3z4",
			want: []string{"3z4", "3Z4"},
		},
		{
			name: "",
			s:    "12345",
			want: []string{"12345"},
		},
		{
			name: "",
			s:    "0",
			want: []string{"0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := letterCasePermutation(tt.s)
			require.Equal(t, tt.want, got)
		})
	}
}
