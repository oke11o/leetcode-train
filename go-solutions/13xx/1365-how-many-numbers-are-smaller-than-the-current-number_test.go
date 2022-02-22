package _3xx

import (
	"reflect"
	"sort"
	"testing"
)

/**
1365. How Many Numbers Are Smaller Than the Current Number
Easy
https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/

Questions:
- Is there only uniq values?

Next:
- https://leetcode.com/problems/count-of-smaller-numbers-after-self/
*/
func smallerNumbersThanCurrent(nums []int) []int {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	sort.Ints(tmp)
	m := make(map[int]int)
	for i, v := range tmp {
		if _, ok := m[v]; !ok {
			m[v] = i
		}
	}
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = m[v]
	}

	return result
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_smallerNumbersThanCurrent(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				nums: []int{8, 1, 2, 2, 3},
			},
			want: []int{4, 0, 1, 1, 3},
		},
		{
			name: "",
			args: args{
				nums: []int{6, 5, 4, 8},
			},
			want: []int{2, 1, 0, 3},
		},
		{
			name: "",
			args: args{
				nums: []int{7, 7, 7, 7},
			},
			want: []int{0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallerNumbersThanCurrent(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallerNumbersThanCurrent() = %v, want %v", got, tt.want)
			}
		})
	}
}
