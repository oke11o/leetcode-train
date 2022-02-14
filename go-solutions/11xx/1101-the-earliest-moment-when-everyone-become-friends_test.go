package _1xx

import (
	"sort"
	"testing"
)

/**
https://leetcode.com/problems/the-earliest-moment-when-everyone-become-friends/
1101. The Earliest Moment When Everyone Become Friends
Medium
[#union_find]
*/
func earliestAcq(logs [][]int, n int) int {
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})
	uf := NewUnionFind(n)
	result := n
	for _, log := range logs {

		if uf.connect(log[1], log[2]) {
			result--
		}

		if result == 1 {
			return log[0]
		}
	}
	return -1
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
	parentA := u.parent(a)
	parentB := u.parent(b)
	if parentA == parentB {
		return false
	}
	if u.rating[parentA] == u.rating[parentB] {
		u.parents[parentA] = b
		u.rating[parentB]++
	} else if u.rating[parentA] < u.rating[parentB] {
		u.parents[parentA] = b
		u.rating[parentB]++
	} else {
		u.parents[parentB] = a
		u.rating[parentA]++
	}

	return true
}
func (u *UnionFind) parent(a int) int {
	if u.parents[a] == a {
		return a
	}
	return u.parent(u.parents[a])
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_earliestAcq(t *testing.T) {
	type args struct {
		logs [][]int
		n    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				logs: [][]int{{20190101, 0, 1}, {20190104, 3, 4}, {20190107, 2, 3}, {20190211, 1, 5}, {20190224, 2, 4}, {20190301, 0, 3}, {20190312, 1, 2}, {20190322, 4, 5}},
				n:    6,
			},
			want: 20190301,
		},
		{
			name: "",
			args: args{
				logs: [][]int{{0, 2, 0}, {1, 0, 1}, {3, 0, 3}, {4, 1, 2}, {7, 3, 1}},
				n:    4,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				logs: [][]int{{5, 4, 3}, {2, 0, 4}, {1, 1, 2}, {0, 0, 2}, {9, 1, 3}, {3, 1, 4}, {8, 2, 4}, {6, 1, 0}},
				n:    5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := earliestAcq(tt.args.logs, tt.args.n); got != tt.want {
				t.Errorf("earliestAcq() = %v, want %v", got, tt.want)
			}
		})
	}
}
