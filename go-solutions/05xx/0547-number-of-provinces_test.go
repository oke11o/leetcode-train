package _5xx

import "testing"

/**
547. Number of Provinces
Medium
https://leetcode.com/problems/number-of-provinces/
*/
func findCircleNum(isConnected [][]int) int {
	uf := NewUnionFind(len(isConnected))
	for n := 0; n < len(isConnected); n++ {
		for m := 0; m < len(isConnected); m++ {
			if isConnected[n][m] == 1 {
				uf.connect(n, m)
			}
		}
	}
	m := make(map[int]struct{})
	for _, v := range uf.parents {
		m[uf.find(v)] = struct{}{}
	}
	return len(m)
}

type UnionFind struct {
	rating  []int
	parents []int
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
func Test_findCircleNum(t *testing.T) {
	type args struct {
		isConnected [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				isConnected: [][]int{
					{1, 1, 0},
					{1, 1, 0},
					{0, 0, 1},
				},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				isConnected: [][]int{
					{1, 0, 0},
					{0, 1, 0},
					{0, 0, 1},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCircleNum(tt.args.isConnected); got != tt.want {
				t.Errorf("findCircleNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
