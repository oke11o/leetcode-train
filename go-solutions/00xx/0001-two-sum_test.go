package _0xx

import (
	"reflect"
	"testing"
)

/**
0001. Two Sum
https://leetcode.com/problems/two-sum/
Easy
[#array, #hash_table]

*/
// Brute force
// TC: O(n^2)
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum_hashTable_twoPass(nums []int, target int) []int {
	hm := make(map[int]int, len(nums))
	for i, v := range nums {
		hm[v] = i
	}
	for i, v := range nums {
		v2 := target - v
		if i2, ok := hm[v2]; ok && i != i2 {
			return []int{i, i2}
		}
	}
	return nil
}

func twoSum_hashTable_onePass(nums []int, target int) []int {
	hm := make(map[int]int, len(nums))
	for i, v := range nums {
		v2 := target - v
		if i2, ok := hm[v2]; ok && i != i2 {
			return []int{i2, i}
		}
		hm[v] = i
	}
	return nil
}

func Test_twoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
			if got := twoSum_hashTable_onePass(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
			if got := twoSum_hashTable_twoPass(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
