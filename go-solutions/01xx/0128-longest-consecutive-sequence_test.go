package _1xx

import "testing"

/**
128. Longest Consecutive Sequence
Medium
https://leetcode.com/problems/longest-consecutive-sequence/
*/
func longestConsecutive(nums []int) int {
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v] = 1
	}
	result := 0
	for _, v := range nums {
		if _, ok := m[v-1]; !ok { // v-1 - проверяем, что v - это начало нашей последовательности. То есть зайдем во внутренний массив только для тех случаем, когда v - это начало последовательности.
			cur := v
			curRes := 0
			for {
				if _, ok := m[cur]; !ok {
					break
				}
				cur++
				curRes++
			}
			if result < curRes {
				result = curRes
			}
		}
	}
	return result
}

func Test_longestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "",
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			name: "",
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
