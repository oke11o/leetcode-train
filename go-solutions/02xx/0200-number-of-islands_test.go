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
			pos := row*colLen + col
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
	m := make(map[int]int)
	for k := range uf.parents {
		_, ok := units[k]
		if ok {
			par := uf.find(k)
			m[par]++
		}
	}

	return len(m)
}

type UnionFind struct {
	rating  []int
	parents []int
}

func NewUnionFind(n int) *UnionFind {
	parents := make([]int, n)
	for i := range parents {
		parents[i] = i
	}
	return &UnionFind{
		rating:  make([]int, n),
		parents: parents,
	}
}

// connect return flag. Is already connected
func (u *UnionFind) connect(a, b int) bool {
	parentA := u.find(a)
	parentB := u.find(b)
	if parentA == parentB {
		return false
	}
	if u.rating[parentA] == u.rating[parentB] {
		u.parents[parentA] = parentB
		u.rating[parentB]++
	} else if u.rating[parentA] < u.rating[parentB] {
		u.parents[parentA] = parentB
	} else {
		u.parents[parentB] = parentA
	}

	return true
}
func (u *UnionFind) find(a int) int {
	if u.parents[a] == a {
		return a
	}
	return u.find(u.parents[a])
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
					{'1', '1', '1', '1', '0'},
					{'1', '1', '0', '1', '0'},
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
