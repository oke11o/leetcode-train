package _2xx

import "testing"

/**
https://leetcode.com/problems/graph-valid-tree/
261. Graph Valid Tree
Medium
*/
func validTree(n int, edges [][]int) bool {
	if len(edges) < n-1 {
		return false
	}
	uf := NewUnionFindForValidateTree(n)
	for _, edge := range edges {
		if !uf.connect(edge[0], edge[1]) {
			return false
		}
	}
	return true
}

type UnionFindForValidateTree struct {
	parents []int
	ratings []int
}

func NewUnionFindForValidateTree(n int) *UnionFindForValidateTree {
	uf := UnionFindForValidateTree{
		parents: make([]int, n),
		ratings: make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parents[i] = i
	}
	return &uf
}

func (uf *UnionFindForValidateTree) connect(a, b int) bool {
	parentA := uf.find(a)
	parentB := uf.find(b)
	if parentA == parentB {
		return false
	}
	if uf.ratings[parentA] == uf.ratings[parentB] {
		uf.parents[parentA] = uf.parents[parentB]
		uf.ratings[parentB]++
	} else if uf.ratings[parentA] > uf.ratings[parentB] {
		uf.parents[parentB] = uf.parents[parentA]
	} else {
		uf.parents[parentA] = uf.parents[parentB]
	}
	return true
}

func (uf *UnionFindForValidateTree) find(a int) int {
	if uf.parents[a] == a {
		return a
	}
	return uf.find(uf.parents[a])
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_validTree(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				n:     5,
				edges: [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}},
			},
			want: true,
		},
		{
			name: "",
			args: args{
				n:     5,
				edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}},
			},
			want: false,
		},
		{
			name: "",
			args: args{
				n:     4,
				edges: [][]int{{0, 1}, {2, 3}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validTree(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("validTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
