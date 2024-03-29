package main

import (
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

type el [2]int

func (e el) Val() int {
	return e[0]
}
func (e el) Less(v item) bool {
	return e.Val() < v.Val()
}

func TestMyPriorityQueue(t *testing.T) {
	h := MyPriorityQueue[el]{}
	ints := []int{7, 29, 42, 18, 18, 12, 11, 14}
	for _, v := range ints {
		h.Insert(el{v, v})
	}
	sort.Slice(ints, func(i, j int) bool {
		return ints[i] > ints[j]
	})
	var got []int
	for range ints {
		v := h.ExtractMax()
		got = append(got, v[0])
	}
	require.Equal(t, ints, got)
}
