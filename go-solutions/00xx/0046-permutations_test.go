package _0xx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/**
 * 0046. Permutations
 * Medium
 * #backtracking
 */
func permute(nums []int) [][]int {
	resultPermute := make([][]int, 0)

	var permuteBacktrack func(current int, nums []int)
	permuteBacktrack = func(current int, nums []int) {
		if current == len(nums) {
			item := make([]int, len(nums))
			copy(item, nums)
			resultPermute = append(resultPermute, item)
			return
		}
		for i := current; i < len(nums); i++ { // 1, 2, 3, 4
			nums[i], nums[current] = nums[current], nums[i]
			permuteBacktrack(current+1, nums)
			nums[i], nums[current] = nums[current], nums[i]
		}
	}

	permuteBacktrack(0, nums)
	return resultPermute
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_permute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "",
			nums: []int{1},
			want: [][]int{{1}},
		},
		{
			name: "",
			nums: []int{2, 1},
			want: [][]int{{2, 1}, {1, 2}},
		},
		{
			name: "",
			nums: []int{3, 2, 1},
			want: [][]int{
				{3, 2, 1}, {3, 1, 2},
				{2, 3, 1}, {2, 1, 3},
				{1, 2, 3}, {1, 3, 2},
			},
		},
		{
			name: "",
			nums: []int{1, 2, 3, 4},
			want: [][]int{
				{1, 2, 3, 4}, {1, 2, 4, 3}, {1, 3, 2, 4}, {1, 3, 4, 2}, {1, 4, 2, 3}, {1, 4, 3, 2},
				{2, 1, 3, 4}, {2, 1, 4, 3}, {2, 3, 1, 4}, {2, 3, 4, 1}, {2, 4, 3, 1}, {2, 4, 1, 3},
				{3, 1, 2, 4}, {3, 1, 4, 2}, {3, 2, 1, 4}, {3, 2, 4, 1}, {3, 4, 1, 2}, {3, 4, 2, 1},
				{4, 1, 2, 3}, {4, 1, 3, 2}, {4, 2, 1, 3}, {4, 2, 3, 1}, {4, 3, 1, 2}, {4, 3, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.nums)
			require.Equal(t, len(tt.want), len(got))
			require.ElementsMatch(t, tt.want, got)
		})
	}
}
