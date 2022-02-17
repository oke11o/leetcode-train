package _6xx

import (
	"reflect"
	"testing"
)

/**
https://leetcode.com/problems/redundant-connection/
684. Redundant Connection
Medium
[#union_find]
*/
func findRedundantConnection(edges [][]int) []int {
	uf := NewUnionFind(len(edges))
	result := []int{}
	for _, edge := range edges {
		if !uf.connect(edge[0]-1, edge[1]-1) {
			result = []int{edge[0], edge[1]}
		}
	}
	return result
}

type UnionFind struct {
	parents []int
	rating  []int
}

func NewUnionFind(n int) *UnionFind {
	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
	}
	return &UnionFind{
		parents: parents,
		rating:  make([]int, n),
	}
}

// connect return flag. Is already connected
func (uf *UnionFind) connect(a, b int) bool {
	parentA := uf.findParent(a)
	parentB := uf.findParent(b)
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

func (uf *UnionFind) findParent(a int) int {
	if uf.parents[a] == a {
		return a
	}
	return uf.findParent(uf.parents[a])
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_findRedundantConnection(t *testing.T) {
	type args struct {
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				edges: [][]int{{1, 2}, {1, 3}, {2, 3}},
			},
			want: []int{2, 3},
		},
		{
			name: "",
			args: args{
				edges: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}},
			},
			want: []int{1, 4},
		},
		{
			name: "",
			args: args{
				edges: [][]int{{1, 5}, {1, 4}, {1, 2}, {2, 3}, {3, 4}},
			},
			want: []int{3, 4},
		},
		{
			name: "",
			args: args{
				edges: [][]int{{1, 2}, {3, 4}, {1, 4}, {2, 3}, {1, 5}, {1, 6}, {5, 6}, {4, 6}},
			},
			want: []int{4, 6},
		},
		{
			name: "",
			args: args{
				edges: [][]int{{1, 2}, {3, 4}, {1, 4}, {2, 3}, {1, 5}, {5, 6}},
			},
			want: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRedundantConnection(tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRedundantConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
