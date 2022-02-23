package _2xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_addOperators(t *testing.T) {
	tests := []struct {
		name   string
		num    string
		target int
		want   []string
	}{
		{
			name:   "",
			num:    "123",
			target: 6,
			want:   []string{"1*2*3", "1+2+3"},
		},
		{
			name:   "",
			num:    "232",
			target: 8,
			want:   []string{"2*3+2", "2+3*2"},
		},
		{
			name:   "",
			num:    "105",
			target: 5,
			want:   []string{"1*0+5", "10-5"},
		},
		{
			name:   "",
			num:    "00",
			target: 0,
			want:   []string{"0*0", "0+0", "0-0"},
		},
		{
			name:   "",
			num:    "3456237490",
			target: 9191,
			want:   []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addOperators(tt.num, tt.target)
			require.Equal(t, tt.want, got)
		})
	}
}
