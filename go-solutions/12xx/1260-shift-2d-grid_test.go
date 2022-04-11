package _2xx

import (
	"reflect"
	"testing"
)

/**
https://leetcode.com/problems/shift-2d-grid/
1260. Shift 2D Grid
Easy

Мои пояснения.
ОЧЕНЬ важно правильно считать newPos. И newRow, newCol.
Так как мы row - это один набор колонок. Поэтому и считать надо от colLen.
*/
func shiftGrid(grid [][]int, k int) [][]int {
	rowLen := len(grid)
	if rowLen == 0 {
		return grid
	}
	colLen := len(grid[0])
	if colLen == 0 {
		return grid
	}
	length := rowLen * colLen
	k = k % length

	tmp := 0
	buf := make([]int, 0, k+1)
	for row := 0; row < rowLen; row++ {
		for col := 0; col < colLen; col++ {
			newPos := (row*colLen + col + k) % length // !!!
			newRow := newPos / colLen                 // !!!
			newCol := newPos % colLen                 // !!!
			buf = append(buf, grid[newRow][newCol])
			if tmp < k {
				grid[newRow][newCol] = grid[row][col]
				tmp++
			} else {
				grid[newRow][newCol] = buf[0]
				buf = buf[1:]
			}
		}
	}

	return grid
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_shiftGrid(t *testing.T) {
	type args struct {
		grid [][]int
		k    int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				grid: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
				k:    1,
			},
			want: [][]int{{9, 1, 2}, {3, 4, 5}, {6, 7, 8}},
		},
		{
			name: "",
			args: args{
				grid: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
				k:    10,
			},
			want: [][]int{{9, 1, 2}, {3, 4, 5}, {6, 7, 8}},
		},
		{
			name: "",
			args: args{
				grid: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
				k:    4,
			},
			want: [][]int{{6, 7, 8}, {9, 1, 2}, {3, 4, 5}},
		},
		{
			name: "",
			args: args{
				grid: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
				k:    6,
			},
			want: [][]int{{7, 8, 9, 10}, {11, 12, 1, 2}, {3, 4, 5, 6}},
		},
		{
			name: "",
			args: args{
				grid: [][]int{{1}, {2}, {3}, {4}, {7}, {6}, {5}},
				k:    23,
			},
			want: [][]int{{6}, {5}, {1}, {2}, {3}, {4}, {7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shiftGrid(tt.args.grid, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shiftGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
