package _1xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

/*
*
161. One Edit Distance
Medium
#string, #two_pointers
*/
func isOneEditDistance(s string, t string) bool {
	if s == t {
		return false
	}
	isOneModified := func(s string, t string) bool {
		// len(s) == len(t)
		for i := 0; i < len(s); i++ {
			if s[i] != t[i] {
				return s[i+1:] == t[i+1:]
			}
		}
		return true
	}
	isOneDelete := func(s string, t string) bool {
		// len(s) > len(t)
		i := 0
		for ; i < len(t); i++ {
			if s[i] != t[i] {
				break
			}
		}
		return s[i+1:] == t[i:]
	}
	if len(s) == len(t) {
		return isOneModified(s, t)
	} else if len(s) > len(t) {
		return isOneDelete(s, t)
	}
	return isOneDelete(t, s)
}

func Test_isOneEditDistance(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "",
			s:    "ab",
			t:    "acb",
			want: true,
		},
		{
			name: "",
			s:    "",
			t:    "",
			want: false,
		},
		{
			name: "",
			s:    "as",
			t:    "as",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isOneEditDistance(tt.s, tt.t)
			require.Equal(t, tt.want, got)
		})
	}
}
