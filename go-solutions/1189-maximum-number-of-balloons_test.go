package go_solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_maxNumberOfBalloons(t *testing.T) {
	tests := []struct {
		name string
		text string
		want int
	}{
		{
			name: "",
			text: "nlaebolko",
			want: 1,
		},
		{
			name: "",
			text: "loonbalxballpoon",
			want: 2,
		},
		{
			name: "",
			text: "leetcode",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxNumberOfBalloons(tt.text)
			require.Equal(t, tt.want, got)
		})
	}
}
