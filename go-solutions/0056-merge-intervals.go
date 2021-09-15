package go_solutions

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := [][]int{intervals[0]}
	for i := 0; i < len(intervals); i++ {
		val := intervals[i]
		lastIdx := len(result) - 1
		if result[lastIdx][1] >= val[0] {
			if result[lastIdx][1] < val[1] {
				result[lastIdx][1] = val[1]
			}
		} else {
			result = append(result, val)
		}
	}
	return result
}
