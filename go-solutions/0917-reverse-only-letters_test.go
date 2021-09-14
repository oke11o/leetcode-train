package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_reverseOnlyLetters(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "",
			input: "ab-cd",
			want:  "dc-ba",
		},

		{
			name:  "",
			input: "a-bC-dEf-ghIj",
			want:  "j-Ih-gfE-dCba",
		},

		{
			name:  "",
			input: "Test1ng-Leet=code-Q!",
			want:  "Qedo1ct-eeLg=ntse-T!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseOnlyLetters(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
