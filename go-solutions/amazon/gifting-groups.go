package amazon

import (
	"io"
)

/**
Amazon gifting groups
*/
func countGroups(related []string, w io.Writer) int32 {
	//for i, v := range related {
	//	if i >= 90 {
	//		fmt.Fprintf(w, "%s\n", v)
	//	}
	//}
	//return 0
	// return int32(len(related))
	uf := NewUnionFind(len(related))
	for i := 0; i < len(related); i++ {
		for j := i + 1; j < len(related); j++ {
			if i == j {
				continue
			}
			cur := related[i][j]
			if cur == '1' {
				uf.connect(i, j)
			}
		}
	}
	m := make(map[int][]int)
	for i, p := range uf.parents {
		m[p] = append(m[p], i)
	}
	return int32(len(m))
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
		u.parents[parentA] = parentB
		u.rating[parentB]++
	} else if u.rating[parentA] < u.rating[parentB] {
		u.parents[parentA] = parentB
		u.rating[parentB]++
	} else {
		u.parents[parentB] = parentA
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
