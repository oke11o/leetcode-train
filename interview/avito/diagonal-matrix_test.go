package avito

import (
	"reflect"
	"testing"
)

func printInvertDiagonals(matrix [][]int) [][]int {
	result := make([][]int, 0, len(matrix)*2-1)
	for i := 0; i < len(matrix)*2-1; i++ {
		var row, col int
		if i < len(matrix) {
			row = i
		} else {
			col = i - len(matrix) + 1
			row = len(matrix) - 1
		}
		res := make([]int, 0)
		for row >= 0 && col < len(matrix) {
			res = append(res, matrix[row][col])
			row--
			col++
		}

		result = append(result, res)
	}

	return result
}

func Test_printInvertDiagonals(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: [][]int{
				{1},
				{4, 2},
				{7, 5, 3},
				{8, 6},
				{9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printInvertDiagonals(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printInvertDiagonals() = %v, want %v", got, tt.want)
			}
		})
	}
}
