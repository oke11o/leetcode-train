package go_solutions

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{{}}
	exists := make(map[int][]int)
	for _, n := range nums {
		length := len(result)
		for i := 0; i < length; i++ {
			if j, ok := exists[n]; ok && findInSlice(i, j) {
				continue
			}
			exists[n] = append(exists[n], i)

			item := result[i][:len(result[i]):len(result[i])]
			item = append(item, n)
			result = append(result, item)
		}
	}
	return result
}

func findInSlice(i int, s []int) bool {
	for _, n := range s {
		if i == n {
			return true
		}
	}
	return false
}
