package _1xx

import (
	"reflect"
	"testing"
)

/**
167. Two Sum II - Input Array Is Sorted
Medium
https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

[]

*/
func twoSum(numbers []int, target int) []int {
	right := len(numbers) - 1
	for left := 0; left < right; left++ {
		for ; right > left; right-- {
			if numbers[left]+numbers[right] == target {
				return []int{left + 1, right + 1}
			}
			if numbers[left]+numbers[right] < target {
				break
			}
		}
	}
	return []int{}
}

func twoSum_twoPointers(numbers []int, target int) []int {
	right := len(numbers) - 1
	left := 0
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}

func Test_twoSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{
			name:    "",
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2},
		},
		{
			name:    "",
			numbers: []int{2, 3, 4},
			target:  6,
			want:    []int{1, 3},
		},
		{
			name:    "",
			numbers: []int{},
			target:  0,
			want:    []int{},
		},
		{
			name:    "",
			numbers: []int{-1, 0},
			target:  -1,
			want:    []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.numbers, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
			if got := twoSum_twoPointers(tt.numbers, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
