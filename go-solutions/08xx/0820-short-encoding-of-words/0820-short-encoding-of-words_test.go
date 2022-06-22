package _820_short_encoding_of_words

import "testing"

func Test_minimumLengthEncoding(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  int
	}{
		{
			name:  "",
			words: []string{"time", "me", "bell"},
			want:  10,
		},
		{
			name:  "",
			words: []string{"t"},
			want:  2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumLengthEncoding(tt.words); got != tt.want {
				t.Errorf("minimumLengthEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}
