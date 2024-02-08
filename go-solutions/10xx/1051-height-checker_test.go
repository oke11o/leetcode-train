package _0xx

import "testing"

/*
1051. Height Checker
https://leetcode.com/problems/height-checker/description/
Это просто сортировка другая
*/
func heightChecker(heights []int) int {
	studentsHeights := [101]int{}

	for _, height := range heights {
		studentsHeights[height]++
	}

	pointer := 0
	result := 0

	for height, count := range studentsHeights {
		for count > 0 {
			if heights[pointer] != height {
				result++
			}
			pointer++
			count--
		}
	}

	return result
}

func Test_heightChecker(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{
			name:    "",
			heights: []int{1, 1, 4, 2, 1, 3},
			want:    3,
		},
		{
			name:    "",
			heights: []int{5, 1, 2, 3, 4},
			want:    5,
		},
		{
			name:    "",
			heights: []int{1, 2, 3, 4, 5},
			want:    1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := heightChecker(tt.heights); got != tt.want {
				t.Errorf("longestStrChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
