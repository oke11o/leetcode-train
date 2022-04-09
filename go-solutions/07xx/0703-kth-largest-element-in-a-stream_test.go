package _7xx

import (
	"sort"
	"testing"
)

/**
https://leetcode.com/problems/kth-largest-element-in-a-stream/
703. Kth Largest Element in a Stream
Easy
*/
type KthLargest struct {
	queue []int
	k     int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	kl := KthLargest{
		queue: make([]int, 0, k),
		k:     k,
	}
	if len(nums) < k {
		kl.queue = append(kl.queue, nums...)
	} else {
		kl.queue = nums[len(nums)-k:]
	}

	return kl
}

func (this *KthLargest) binSearch(val int) int {
	left := -1
	right := len(this.queue)
	for right-left > 1 {
		i := (right + left) / 2
		if this.queue[i] > val {
			right = i
		} else {
			left = i
		}
	}
	return left
}

func (this *KthLargest) Add(val int) int {
	i := this.binSearch(val)
	if len(this.queue) < this.k {
		if i == -1 {
			this.queue = append([]int{val}, this.queue...)
		} else {
			this.queue = append(this.queue[:i+1], this.queue[i:len(this.queue)]...)
			this.queue[i] = val
		}
	} else {
		this.queue = append(this.queue[1:i+1], this.queue[i:]...)
		this.queue[i] = val
	}

	return this.queue[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kl := Constructor(tt.init.k, tt.init.arr)
			for _, w := range tt.want {
				if got := kl.Add(w[0]); got != w[1] {
					t.Errorf("Add() = %v, want %v", w[0], w[1])
				}
			}
		})
	}
}
