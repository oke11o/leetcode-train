package _8xx

import (
	"reflect"
	"testing"
)

/**
https://leetcode.com/problems/transpose-matrix/
867. Transpose Matrix
Easy
*/
func transpose(matrix [][]int) [][]int {
	result := make([][]int, len(matrix[0]))
	for i := 0; i < len(result); i++ {
		result[i] = make([]int, len(matrix))
	}
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			result[col][row] = matrix[row][col]
		}
	}
	return result
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_transpose(t *testing.T) {
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
				matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			want: [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}},
		},
		{
			name: "",
			args: args{
				matrix: [][]int{{1, 2, 3}, {4, 5, 6}},
			},
			want: [][]int{{1, 4}, {2, 5}, {3, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transpose(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}
