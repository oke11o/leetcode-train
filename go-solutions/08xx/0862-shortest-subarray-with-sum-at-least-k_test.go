package _8xx

import "testing"

/**
862. Shortest Subarray with Sum at Least K
Hard
https://leetcode.com/problems/shortest-subarray-with-sum-at-least-k/
[#sliding_window]

At first you should solve:
https://leetcode.com/problems/minimum-size-subarray-sum/
*/
func shortestSubarray(nums []int, k int) int {
	sum := 0
	localRes := 0
	result := 999999999999
	first, last := 0, 0
	flag := false
	//             f
	//      2, 0, -1, 3, 0, 1, 2, 2
	//      l
	//sum:  2, 2
	//lRes  1, 2
	for first < len(nums) {
		localRes++         //1,2,3,4
		sum += nums[first] //2,2,1,4
		for sum >= k {
			if result > localRes {
				result = localRes //0,0,0,4
			}
			flag = true
			sum -= nums[last] //2
			last++
			localRes-- //3
		}
		first++
	}
	if !flag {
		return -1
	}
	return result
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_shortestSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "",
			nums: []int{84, -37, 32, 40, 95},
			k:    167,
			want: 3,
		},
		{
			name: "",
			nums: []int{2, 0, -1, 3, 0, 1, 2, 2},
			k:    4,
			want: 2,
		},
		{
			name: "",
			nums: []int{1},
			k:    1,
			want: 1,
		},
		{
			name: "",
			nums: []int{1, 2},
			k:    4,
			want: -1,
		},
		{
			name: "",
			nums: []int{2, -1, 2},
			k:    3,
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestSubarray(tt.nums, tt.k); got != tt.want {
				t.Errorf("shortestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
