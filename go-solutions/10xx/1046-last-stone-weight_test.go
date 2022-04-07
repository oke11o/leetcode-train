package _0xx

import (
	"sort"
	"testing"
)

/**
https://leetcode.com/problems/last-stone-weight/
1046. Last Stone Weight
Easy
#binary_search, #sort
*/
func lastStoneWeight(stones []int) int {
	var binarySearch = func(arr []int, x int) int {
		left := -1
		right := len(arr)
		for (right - left) > 1 {
			i := (left + right) / 2
			if arr[i] <= x {
				left = i
			} else {
				right = i
			}
		}
		if left == -1 {
			return 0
		}

		return left
	}

	sort.Ints(stones)
	for len(stones) > 1 {
		x := stones[len(stones)-1]
		y := stones[len(stones)-2]
		if x != y {
			if x > y {
				x -= y
			} else {
				x = y - x
			}
			stones = stones[:len(stones)-1]

			index := binarySearch(stones, x)
			if index == len(stones)-1 {
				stones[len(stones)-1] = x
			} else {
				stones = append(stones[:index+1], stones[index:len(stones)-1]...)
				if stones[index+1] < x { // TODO: need check and fix. I dont understand WHY?!!!
					stones[index+1] = x
				} else {
					stones[index] = x
				}
			}
		} else {
			stones = stones[:len(stones)-2]
		}
	}
	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_lastStoneWeight(t *testing.T) {
	type args struct {
		stones []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				stones: []int{2, 7, 4, 1, 8, 1},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				stones: []int{1},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				stones: []int{7, 6, 7, 6, 9},
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				stones: []int{316, 157, 73, 106, 771, 828, 46, 212, 926, 604, 600, 992, 71, 51, 477, 869, 425, 405, 859, 924, 45, 187, 283, 590, 303, 66, 508, 982, 464, 398},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastStoneWeight(tt.args.stones); got != tt.want {
				t.Errorf("lastStoneWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
