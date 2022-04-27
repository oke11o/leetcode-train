package _2xx

import (
	"sort"
	"testing"
)

/**
https://leetcode.com/problems/smallest-string-with-swaps/
1202. Smallest String With Swaps
Medium
*/
func smallestStringWithSwaps(s string, pairs [][]int) string {
	uf := NewUnionFind(len(s))
	for _, pair := range pairs {
		uf.connect(pair[0], pair[1])
	}

	// Group the chars that are in the same component
	groups := make(map[int][]rune)
	indeces := make(map[int][]int)
	for i := 0; i < len(s); i++ {
		parent := uf.find(i)
		indeces[parent] = append(indeces[parent], i)
		groups[parent] = append(groups[parent], rune(s[i]))
	}

	result := make([]rune, len(s))
	for parent, group := range groups {
		sort.Slice(group, func(i, j int) bool {
			return group[i] < group[j]
		})
		for i, idx := range indeces[parent] {
			result[idx] = group[i]
		}
	}

	return string(result)
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
	for i := 0; i < n; i++ {
		uf.parents[i] = i
	}
	return uf
}

// return true if it is new connection. Return false if already connected
func (uf *UnionFind) connect(a, b int) bool {
	pa := uf.find(a)
	pb := uf.find(b)
	if pa == pb {
		return false
	}
	if uf.rating[pa] == uf.rating[pb] {
		uf.parents[pa] = pb
		uf.rating[pb]++
	} else if uf.rating[pa] < uf.rating[pb] {
		uf.parents[pa] = pb
	} else {
		uf.parents[pb] = pa
	}
	return true
}

func (uf *UnionFind) find(a int) int {
	if a == uf.parents[a] {
		return a
	}
	return uf.find(uf.parents[a])
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_smallestStringWithSwaps(t *testing.T) {
	type args struct {
		s     string
		pairs [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s:     "dcab",
				pairs: [][]int{{0, 3}, {1, 2}},
			},
			want: "bacd",
		},
		{
			name: "",
			args: args{
				s:     "dcab",
				pairs: [][]int{{0, 3}, {1, 2}, {0, 2}},
			},
			want: "abcd",
		},
		{
			name: "",
			args: args{
				s:     "cba",
				pairs: [][]int{{0, 1}, {1, 2}},
			},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestStringWithSwaps(tt.args.s, tt.args.pairs); got != tt.want {
				t.Errorf("smallestStringWithSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
