package _3xx

import "testing"

/**
https://leetcode.com/problems/maximum-size-subarray-sum-equals-k/
325. Maximum Size Subarray Sum Equals k
Medium
*/

/**
If we run into a duplicate (which is possible because of negative numbers), we should not update the index in the hash map because we want the longest subarray, so we want to keep the index as far to the left as possible.
*/
func maxSubArrayLen(nums []int, k int) int {
	result := 0
	prefixSum := 0
	sums := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		prefixSum += nums[i]

		// Check if all of the numbers seen so far sum to k.
		if prefixSum == k {
			result = i + 1
		}

		// If any subarray seen so far sums to k, then
		// update the length of the longest_subarray.
		if v, ok := sums[prefixSum-k]; ok {
			if result < i-v {
				result = i - v
			}
		}

		// Only add the current prefix_sum index pair to the
		// map if the prefix_sum is not already in the map.
		if _, ok := sums[prefixSum]; !ok {
			sums[prefixSum] = i
		}
	}
	return result
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_maxSubArrayLen(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "",
			nums: []int{1, -1, 5, -2, 3},
			k:    3,
			want: 4,
		},
		{
			name: "",
			nums: []int{-2, -1, 2, 1},
			k:    1,
			want: 2,
		},
		{
			name: "",
			nums: []int{1, 0, -1},
			k:    -1,
			want: 2,
		},
		{
			name: "",
			nums: []int{},
			k:    0,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArrayLen(tt.nums, tt.k); got != tt.want {
				t.Errorf("maxSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
