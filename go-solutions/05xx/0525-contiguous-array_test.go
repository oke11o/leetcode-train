package _5xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// 525. Contiguous Array
// https://leetcode.com/problems/contiguous-array/description/
/**
Брутфорсом тоже не просто решить. Но становится проще, если подумать, что мы можем посчитать за один проход от 0...len?
Можем считать кол-во нулей и единиц. И если оно равно, сравнивать с предыдущим максимумом.
А потом сдвигаемся на 1...len.
Это и есть брутфорс

*/
func findMaxLength_brut(nums []int) int {
	answer := 0
	for i := 0; i < len(nums); i++ {
		zeros, units := 0, 0
		for j := i; j < len(nums); j++ {
			if nums[j] == 0 {
				zeros++
			} else {
				units++
			}
			if zeros == units {
				// можно считать отдельно счетчик, сколько прошли. Но на самом деле это j, но не с начала, а от i то есть
				r := j - i + 1 //
				answer = max(answer, r)
			}
		}
	}
	return answer
}

func Test_findMaxLength(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "",
			nums: []int{0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0},
			want: 10,
		},
		{
			name: "",
			nums: []int{0, 1},
			want: 2,
		},
		{
			name: "",
			nums: []int{0, 1, 0},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, findMaxLength_brut(tt.nums))
		})
	}
}
