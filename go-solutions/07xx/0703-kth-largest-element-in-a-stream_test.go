package _7xx

import (
	"container/heap"
	"fmt"
	"testing"
)

/**
https://leetcode.com/problems/kth-largest-element-in-a-stream/
703. Kth Largest Element in a Stream
Easy
*/
type KthLargest struct {
	heap *IntHeap
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	h := IntHeap(nums)
	kl := KthLargest{
		heap: &h,
		k:    k,
	}
	heap.Init(kl.heap)
	for kl.heap.Len() > k {
		heap.Pop(kl.heap)
	}

	return kl
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.heap, val)
	if this.heap.Len() > this.k {
		heap.Pop(this.heap)
	}

	return (*this.heap)[0]
}

/**
HEAP
*/
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/*********************************/
/************* TESTS *************/
/*********************************/
func TestHead(t *testing.T) {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	tests := []int{1, 2, 3, 5}
	for i, want := range tests {
		l := len(tests) - i
		if l != h.Len() {
			t.Errorf("Len() = %v, want %v", h.Len(), l)
		}
		if got := heap.Pop(h); got != want {
			t.Errorf("Pop() = %v, want %v", got, want)
		}
	}
}

/*********************************/
/************* TESTS *************/
/*********************************/
func TestKthLargest_Add(t *testing.T) {
	type init struct {
		k   int
		arr []int
	}
	tests := []struct {
		name string
		init init
		want [][2]int
	}{
		{
			name: "",
			init: init{
				k:   3,
				arr: []int{4, 5, 8, 2},
			},
			want: [][2]int{
				{3, 4},
				{5, 5},
				{10, 5},
				{9, 8},
				{4, 8},
			},
		},
		{
			name: "",
			init: init{
				k:   1,
				arr: []int{},
			},
			want: [][2]int{
				{-3, -3},
				{-2, -2},
				{-4, -2},
				{0, 0},
				{4, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kl := Constructor(tt.init.k, tt.init.arr)
			for _, w := range tt.want {
				if got := kl.Add(w[0]); got != w[1] {
					t.Errorf("Add() = %v, want %v", got, w[1])
				}
			}
		})
	}
}
