package _2xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_gameOfLife(t *testing.T) {
	type args struct {
		board [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				board: [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}},
			},
			want: [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameOfLife(tt.args.board)
			require.Equal(t, tt.want, tt.args.board)
		})
	}
}
