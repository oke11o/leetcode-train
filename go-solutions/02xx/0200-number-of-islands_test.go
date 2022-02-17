package _2xx

import "testing"

/**
200. Number of Islands
Medium
https://leetcode.com/problems/number-of-islands/
*/
func numIslands(grid [][]byte) int {
	rowLen := len(grid)
	colLen := len(grid[0])
	uf := NewUnionFind(colLen * rowLen)
	units := make(map[int]struct{}, colLen*rowLen)
	for row := 0; row < rowLen; row++ {
		for col := 0; col < colLen; col++ {
			pos := row*colLen + col // !!! Важно правильно вычислять позицию
			if grid[row][col] == '1' {
				units[pos] = struct{}{}
			}
			if col+1 < colLen {
				if grid[row][col] == '1' && grid[row][col] == grid[row][col+1] {
					rightPos := row*colLen + col + 1
					uf.connect(pos, rightPos)
				}
			}
			if row+1 < rowLen {
				if grid[row][col] == '1' && grid[row][col] == grid[row+1][col] {
					downPos := (row+1)*colLen + col
					uf.connect(pos, downPos)
				}
			}
		}
	}
	// calc zeroes island count
	m := make(map[int]struct{})
	for k := range uf.parents {
		_, ok := units[k]
		if ok {
			par := uf.find(k)
			m[par] = struct{}{}
		}
	}

	return len(m)
}

type UnionFind struct {
	unitCount int
	rating    []int
	parents   []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		rating:  make([]int, n),
		parents: make([]int, n),
	}
	for i := range uf.parents {
		uf.parents[i] = i
	}
	return uf
}

// connect return flag. Is already connected
func (uf *UnionFind) connect(a, b int) bool {
	parentA := uf.find(a)
	parentB := uf.find(b)
	if parentA == parentB {
		return false
	}
	if uf.rating[parentA] == uf.rating[parentB] {
		uf.parents[parentA] = parentB
		uf.rating[parentB]++
	} else if uf.rating[parentA] < uf.rating[parentB] {
		uf.parents[parentA] = parentB
	} else {
		uf.parents[parentB] = parentA
	}
	uf.unitCount--

	return true
}

func (uf *UnionFind) find(a int) int {
	if uf.parents[a] == a {
		return a
	}
	return uf.find(uf.parents[a])
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_numIslands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				grid: [][]byte{
					{'1', '1'},
				},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				grid: [][]byte{
					{'1'},
					{'1'},
				},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				grid: [][]byte{
					{'1', '1', '0'},
					{'1', '1', '0'},
					{'0', '0', '1'},
				},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				grid: [][]byte{
					// unitsLen == 9
					{'1', '1', '1', '1', '0'}, //2,4,5,6
					{'1', '1', '0', '1', '0'}, //7,8
					{'1', '1', '0', '0', '0'},
					{'0', '0', '0', '0', '0'},
				},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				grid: [][]byte{
					{'1', '1', '0', '0', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '1', '0', '0'},
					{'0', '0', '0', '1', '1'},
				},
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				grid: [][]byte{
					{'1', '0', '1', '1', '1'},
					{'1', '0', '1', '0', '1'},
					{'1', '1', '1', '0', '1'},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands(tt.args.grid); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
