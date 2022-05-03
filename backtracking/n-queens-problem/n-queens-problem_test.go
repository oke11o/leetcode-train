package n_queens_problem

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_queensProblems(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]int
	}{
		{
			name: "",
			n:    4,
			want: [][]int{{1, 3, 0, 2}, {2, 0, 3, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := queensProblems(tt.n)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_boundFunc(t *testing.T) {
	tests := []struct {
		name    string
		row     int
		col     int
		current []int
		want    bool
	}{
		{
			name:    "",
			row:     2,
			col:     0,
			current: []int{1, 1},
			want:    false,
		},
		{
			name:    "",
			row:     2,
			col:     1,
			current: []int{1, 1},
			want:    false,
		},
		{
			name:    "",
			row:     2,
			col:     2,
			current: []int{1, 1},
			want:    false,
		},
		{
			name:    "",
			row:     2,
			col:     3,
			current: []int{1, 1},
			want:    false,
		},
		{
			name:    "",
			row:     2,
			col:     4,
			current: []int{1, 1},
			want:    true,
		},
		{
			name:    "",
			row:     2,
			col:     0,
			current: []int{3, 3},
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := boundFunc(tt.row, tt.col, tt.current)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_convert(t *testing.T) {
	tests := []struct {
		name string
		in   [][]int
		want [][]string
	}{
		{
			name: "",
			in:   [][]int{{1, 3, 0, 2}},
			want: [][]string{{".Q..", "...Q", "Q...", "..Q."}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
