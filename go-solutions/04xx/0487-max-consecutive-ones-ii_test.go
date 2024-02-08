package _4xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

/*
*
https://leetcode.com/problems/max-consecutive-ones-ii/
487. Max Consecutive Ones II
Medium
#sliding_window
*/
func findMaxConsecutiveOnes_ii(nums []int) int {
	var ans int

	k := 1
	zeroes := 0
	units := 0

	first := 0
	second := 0
	//L_
	//[1,0,1,1,0]
	//F  ^
	//z=1
	//u=1
	//a=1
	for first < len(nums) {
		val1 := nums[first] // 0
		first++             // 2
		units++             // 2
		if val1 == 0 {
			zeroes++ //1
		}
		for zeroes > k { //false
			val2 := nums[second]
			second++
			units--
			if val2 == 0 {
				zeroes--
			}
		}
		if ans < units {
			ans = units
		}
	}

	return ans
}

func Test_findMaxConsecutiveOnes_ii(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "",
			nums: []int{1, 0, 1, 1, 0},
			want: 4,
		},
		{
			name: "",
			nums: []int{1, 0, 1, 1, 0, 1},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxConsecutiveOnes_ii(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
