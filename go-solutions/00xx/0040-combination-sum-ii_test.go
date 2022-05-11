package _0xx

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
https://leetcode.com/problems/combination-sum-ii/submissions/
40. Combination Sum II
Medium
TODO: not working. Need check solution
*/
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}

	var backtrack func(pos int, curDigits []int, curSum int)
	backtrack = func(pos int, curDigits []int, curSum int) {
		if curSum == target {
			tmp := make([]int, len(curDigits))
			copy(tmp, curDigits)
			for _, r := range result {
				if len(r) != len(curDigits) {
					continue
				}
				flag := true
				for i := 0; i < len(r); i++ {
					if r[i] != curDigits[i] {
						flag = false
						break
					}
				}
				if flag {
					return
				}
			}
			result = append(result, tmp)
			return
		}
		digits := make([]int, len(curDigits)+1)
		copy(digits, curDigits)
		for i := pos; i < len(candidates); i++ {
			v := candidates[i]
			digits[len(digits)-1] = v
			newSum := curSum + v
			if newSum > target {
				return
			}
			backtrack(i+1, digits, newSum)
		}
	}
	backtrack(0, []int{}, 0)

	return result
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_combinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				candidates: []int{10, 1, 2, 7, 6, 1, 5},
				target:     8,
			},
			want: [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
		},
		{
			name: "",
			args: args{
				candidates: []int{2, 5, 2, 1, 2},
				target:     5,
			},
			want: [][]int{{1, 2, 2}, {5}},
		},
		{
			name: "Time limit",
			args: args{
				candidates: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				target:     30,
			},
			want: [][]int{{1, 2, 2}, {5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, combinationSum2(tt.args.candidates, tt.args.target), "combinationSum2(%v, %v)", tt.args.candidates, tt.args.target)
		})
	}
}
